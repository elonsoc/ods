package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"

	locations "github.com/elonsoc/center/backend/locations"
	"github.com/elonsoc/center/backend/service"
)

// IdentityKeys are different from API keys in that they are used to validate the identity of the user
// and are used to determine what data the user can access.
// this map is just a proof of concept, in the future we would swap this out and call our database.
var IdentityKeys = map[string]string{
	"elon_ods:12345": "elon",
}

// CheckAuth is a custom middleware that checks the Authorization header for a valid API key
func CheckAPIKey() func(next http.Handler) http.Handler {
	// this mocks a database of API keys
	APIKEYS := map[string]bool{
		"elon_ods:12345": true,
	}

	// this middleware is a proof of concept, in the future we would swap this out
	// for a call to our database to verify the API key
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// this is where the call we'd make to the database to verify the API key
			// would happen. For now, we just check if the API key is in the map above.
			if !APIKEYS[auth] {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// if both are okay, then we're free to further this request.
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// func(next http.Handler) http.Handler

func CheckIdentity() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token_cookie, err := r.Cookie("identity")
			if err != nil {
				// may log this?
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auth := token_cookie.Value
			if auth == "" {

				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if IdentityKeys[auth] == "" {

				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}

func CustomLogger(log service.LoggerIFace, stat service.StatIFace) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// pass along the http request before we log it
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			log.InfoWithFields(logrus.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"request_id": middleware.GetReqID(r.Context()),
				"ip":         r.RemoteAddr,
				"scheme":     scheme,
				"status":     ww.Status(),
			}, "Request received")
			stat.Increment("request", statsd.IntTag("status", ww.Status()), statsd.StringTag("path", r.URL.Path))
		}

		return http.HandlerFunc(fn)
	}
}

// initialize begins the startup process for the backend of ods.
// At the beginning, we create a new instance of the router, declare usage of multiple middlewares
// initialize connections to external services, and mount the various routers for the apis that we
// will be serving.
func initialize(servicePort, databaseURL, redisURL, loggingURL, statsdURL string) chi.Router {
	startInitialization := time.Now()
	// This is where we initialize the various services that we will be using
	// like the database, logger, stats, etc.
	// in the future, we could split up the apis into separate services
	// which would allow for more flexibility in scaling, allow for more granular control over the API keys,
	// and allow for more granular control over the services that are running.
	// This particular technique is called dependency injection, and it's a good practice to use
	// when writing code that could one day be decoupled into separate services.
	// There are better ways to do this, but this is a good start to keep the app monolithic for now.
	Services := service.NewService(loggingURL, databaseURL, statsdURL)

	// get port from environment variable

	// Create a new instance of the router
	r := chi.NewRouter()

	// Middleware is defined next and works in a LIFO (Last In First Out) order.
	// This means that the first middleware that is declared will be the last one to be executed.

	// middleware.Logger prints a log line for each request (access log)
	r.Use(CustomLogger(Services.Log, Services.Stat))
	r.Use(middleware.RequestID)
	// middleware.RealIP is used to get the real IP address of the client
	r.Use(middleware.RealIP)
	// middleware.Recoverer recovers from panics without crashing the server and supplies
	// a 500 error page (Internal Server Error)
	r.Use(middleware.Recoverer)

	// This is where we mount the various routers for the various APIs that we will be serving
	// again, in the future we can split these up into separate services if we want to for the above reasons
	// routers are versioned, so we can have multiple versions of the same API running at the same time
	// which is a good practice for API design to allow for backwards compatibility and support for
	// older clients who may rely on older versions of the API.

	// this group is for the API that will be used by applications to access the data
	r.Group(func(r chi.Router) {
		// This custom middleware checks the Authorization header for a valid API key
		// and if it's not valid, it returns a 401 Unauthorized error
		r.Use(CheckAPIKey())

		// This get request is just a simple ping endpoint to test that the server is running
		// and that the API key is valid.
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		r.Mount("/locations", locations.NewLocationsRouter(&locations.LocationsRouter{Svcs: Services}).Router)
	})

	// this group is for the API that will be used by the frontend to validate the user's identity
	r.Mount("/identity", r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.AllowContentType("application/json"))
		// Here, we take the token that is passed in the request body and validate it
		// by making a call to the identity service.
		// If the token is valid, we commit this to memory and use it to verify
		// actions that the user takes on the website.
		// So, it works like this:
		// 1. User clicks "Log In" and is redirected to the identity service
		// 2. User logs in and is redirected back to the frontend with a token, we call this a dirty token
		// because it's not yet validated and has been touched by the user
		// 3. We make a call to the identity service to validate the token, and if it's valid,
		// we commit a new token pair to memory, one that is clean (only on the server) and one that is dirty (on the client)
		// the dirty token is a JWT that is used to verify the user's identity and is stored in a cookie.
		// the clean token, instead, is durable and is stored in on our side for a longer period of time.
		// This allows us to verify the user's identity without having to make a call to the identity service every time
		r.Post("/validate", func(w http.ResponseWriter, r *http.Request) {
			type request struct {
				Token string `json:"token"`
			}

			var req request
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				fmt.Println("Error decoding request body: ", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if req.Token == "" {
				fmt.Println("Token is empty")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resp, err := http.Get("http://localhost:1338/validate?token=" + req.Token)
			type validationResponse struct {
				Token string `json:"token"`
			}
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			if resp.StatusCode != 200 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// read the response body into validationresponse
			res := validationResponse{}
			err = json.NewDecoder(resp.Body).Decode(&res)
			if err != nil {
				fmt.Println("Error decoding response body: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// if the token is valid, we commit it to memory

			IdentityKeys[res.Token] = "elon_ods:12345"
			w.Write([]byte("elon_ods:12345"))
		})
	}))

	r.Mount("/affiliate", r.Group(func(r chi.Router) {
		r.Use(CheckIdentity())

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("you're a true affiliate."))
		})
	}))

	Services.Log.Info("Server running on port", servicePort)
	Services.Stat.TimeElapsed("server.start", time.Since(startInitialization).Milliseconds())
	return r
}

func main() {
	// get our pertinent information from the environment variables or the command line
	servicePort := flag.String("port", os.Getenv("PORT"), "port to run server on")
	databaseURL := flag.String("database_url", os.Getenv("DATABASE_URL"), "database url")
	redisURL := flag.String("redis_url", os.Getenv("REDIS_URL"), "redis url")
	loggingURL := flag.String("logging_url", os.Getenv("LOGGING_URL"), "logging url")
	statsdURL := flag.String("statsd_url", os.Getenv("STATSD_URL"), "statsd url")
	flag.Parse()
	if *servicePort == "" {
		log.Fatal("port not set")
	}
	if *databaseURL == "" {
		log.Fatal("database url not set")
	}
	if *redisURL == "" {
		log.Fatal("redis url not set")
	}
	if *loggingURL == "" {
		log.Fatal("logging url not set")
	}
	if *statsdURL == "" {
		log.Fatal("statsd url not set")
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", *servicePort),
		initialize(*servicePort, *databaseURL, *redisURL, *loggingURL, *statsdURL))
	if err != nil {
		fmt.Println(err)
	}
}
