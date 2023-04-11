package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/crewjam/saml/samlsp"
	locations "github.com/elonsoc/center/backend/locations"
	"github.com/elonsoc/center/backend/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/smira/go-statsd"
)

// IdentityKeys are different from API keys in that they are used to validate the identity of the user
// and are used to determine what data the user can access.
// this map is just a proof of concept, in the future we would swap this out and call our database.
var IdentityKeys = map[string]string{
	"elon_ods:12345": "elon",
}

// CheckAuth is a custom middleware that checks the Authorization header for a valid API key
func CheckAPIKey(log *logrus.Logger) func(next http.Handler) http.Handler {
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

func CheckIdentity(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log.Info("Checking identity")
			token_cookie, err := r.Cookie("identity")
			if err != nil {
				log.Error(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auth := token_cookie.Value
			if auth == "" {
				log.Error("no auth")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if IdentityKeys[auth] == "" {
				log.Println("no identity key associated with token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}

func CustomLogger(log *logrus.Logger, stat *statsd.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// pass along the http request before we log it
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			scheme := "http"

			if r.TLS != nil {
				scheme = "https"
			}

			log.WithFields(logrus.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"request_id": middleware.GetReqID(r.Context()),
				"ip":         r.RemoteAddr,
				"scheme":     scheme,
				"status":     ww.Status(),
			}).Info("Request received")
			stat.Incr("request", 1, statsd.IntTag("status", ww.Status()), statsd.StringTag("path", r.URL.Path))
		}

		return http.HandlerFunc(fn)
	}
}

// initialize begins the startup process for the backend of ods.
// At the beginning, we create a new instance of the router, declare usage of multiple middlewares
// initialize connections to external services, and mount the various routers for the apis that we
// will be serving.
func initialize(servicePort, databaseURL, redisURL, loggingURL, statsdURL string) chi.Router {
	// This is where we initialize the various services that we will be using
	// like the database, logger, stats, etc.
	// in the future, we could split up the apis into separate services
	// which would allow for more flexibility in scaling, allow for more granular control over the API keys,
	// and allow for more granular control over the services that are running.
	// This particular technique is called dependency injection, and it's a good practice to use
	// when writing code that could one day be decoupled into separate services.
	// There are better ways to do this, but this is a good start to keep the app monolithic for now.
	svc := service.NewService(loggingURL, databaseURL, statsdURL)

	// Create a new instance of the router
	r := chi.NewRouter()

	// Middleware is defined next and works in a LIFO (Last In First Out) order.
	// This means that the first middleware that is declared will be the last one to be executed.

	// middleware.Logger prints a log line for each request (access log)
	r.Use(CustomLogger(svc.Logger, svc.Stat))
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

	// this saml subpath handles the necessary get and post operations to support SAML authentication
	r.Mount("/saml", r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/metadata", svc.Saml.ServeMetadata)
		r.Post("/acs", svc.Saml.ServeACS)
	}))

	// this group is for the API that will be used by applications to access the data
	r.Group(func(r chi.Router) {
		// This custom middleware checks the Authorization header for a valid API key
		// and if it's not valid, it returns a 401 Unauthorized error
		r.Use(CheckAPIKey(svc.Logger))

		// This get request is just a simple ping endpoint to test that the server is running
		// and that the API key is valid.
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		r.Mount("/locations", locations.NewLocationsRouter(&locations.LocationsRouter{Svcs: svc}).Router)
	})

	// This represents endpoints that humans will access with their browser and thus need to affiliate themselves with Elon University
	r.Mount("/affiliate", r.Group(func(r chi.Router) {
		r.Use(svc.Saml.RequireAccount)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", samlsp.AttributeFromContext(r.Context(), "displayName"))
			w.Write([]byte("you're a true affiliate."))
		})
	}))

	svc.Logger.Info("Server running on port %s", servicePort)
	return r
}

func main() {
	// get our pertinent information from the environment variables or the command line
	servicePort := flag.String("port", os.Getenv("PORT"), "port to run server on")
	databaseURL := flag.String("database_url", os.Getenv("DATABASE_URL"), "database url")
	redisURL := flag.String("redis_url", os.Getenv("REDIS_URL"), "redis url")
	loggingURL := flag.String("logging_url", os.Getenv("LOGGING_URL"), "logging url")
	statsdURL := flag.String("statsd_url", os.Getenv("STATSD_URL"), "statsd url")

	// this could use some improvement in nameing and probably would require
	// Hashicorp Vault or someting of the sort
	// x509KeyPair := flag.String("keypair_location", os.Getenv("X509_Keypair_Location"), "location of x509 key pair")

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

	err := http.ListenAndServe(fmt.Sprintf(":%s", *servicePort), initialize(*servicePort, *databaseURL, *redisURL, *loggingURL, *statsdURL))
	if err != nil {
		fmt.Println(err)
	}
}
