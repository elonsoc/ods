package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/elonsoc/saml/samlsp"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"

	"github.com/elonsoc/ods/backend/applications"
	locations "github.com/elonsoc/ods/backend/locations"
	"github.com/elonsoc/ods/backend/service"
)

// CheckAuth is a custom middleware that checks the Authorization header for a valid API key
func CheckAPIKey(db service.DbIFace, stat service.StatIFace) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// this is where the call we'd make to the database to verify the API key
			// would happen. For now, we just check if the API key is in the map above.
			if !db.IsValidApiKey(auth) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// if both are okay, then we're free to further this request.
			stat.Increment("api_key_used", statsd.StringTag("api_key", auth))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func CheckIdentity(tokenSvcr service.TokenIFace, log service.LoggerIFace) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token_cookie, err := r.Cookie("ods_login_cookie_nomnom")
			if err != nil {
				// may log this?
				log.Info("login cookie not presented", nil)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auth := token_cookie.Value
			if auth == "" {
				log.Info("value not available in login cookie", nil)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			res, err := tokenSvcr.ValidateToken(auth)
			if !res || err != nil {
				log.Info("login cookie invalid.", nil)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// getDomainFromURI formats a domain string into a proper
// domain name to be inlayed into a cookie.
func getDomainFromURI(domain string) (string, error) {
	if strings.ToLower(domain[:4]) == "http" {
		u, err := url.Parse(domain)
		if err != nil {
			return "", err
		}
		return u.Hostname(), nil
	}
	// the provided domain is not a URL, so it should be a hostname
	domain, _, err := net.SplitHostPort(domain)
	if err != nil {
		return "", err

	}

	return domain, nil

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

			log.Info("", logrus.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"request_id": middleware.GetReqID(r.Context()),
				"ip":         r.RemoteAddr,
				"scheme":     scheme,
				"status":     ww.Status(),
			})
			stat.Increment("request", statsd.IntTag("status", ww.Status()), statsd.StringTag("path", r.URL.Path))
		}

		return http.HandlerFunc(fn)
	}
}

// initialize begins the startup process for the backend of ods.
// At the beginning, we create a new instance of the router, declare usage of multiple middlewares
// initialize connections to external services, and mount the various routers for the apis that we
// will be serving.
func initialize(servicePort, databaseURL, redisURL, loggingURL, statsdURL, certPath, keyPath, idpURL, spURL, webURL string) chi.Router {
	startInitialization := time.Now()
	// This is where we initialize the various services that we will be using
	// like the database, logger, stats, etc.
	// in the future, we could split up the apis into separate services
	// which would allow for more flexibility in scaling, allow for more granular control over the API keys,
	// and allow for more granular control over the services that are running.
	// This particular technique is called dependency injection, and it's a good practice to use
	// when writing code that could one day be decoupled into separate services.
	// There are better ways to do this, but this is a good start to keep the app monolithic for now.
	svc := service.NewService(loggingURL, databaseURL, statsdURL, certPath, keyPath, idpURL, spURL, webURL)

	samlMiddleware := svc.Saml.GetSamlMiddleware()

	// Create a new instance of the router
	r := chi.NewRouter()

	// Middleware is defined next and works in a LIFO (Last In First Out) order.
	// This means that the first middleware that is declared will be the last one to be executed.

	// middleware.Logger prints a log line for each request (access log)
	r.Use(CustomLogger(svc.Log, svc.Stat))
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
	// r.Mount("/saml", r.Group(func(r chi.Router) {
	// 	r.Use(middleware.NoCache)

	// 	r.Get("/metadata", samlMiddleware.ServeMetadata)
	// 	r.Post("/acs", samlMiddleware.ServeACS)
	// }))

	// this group is for the API that will be used by applications to access the data
	r.Group(func(r chi.Router) {
		// This custom middleware checks the Authorization header for a valid API key
		// and if it's not valid, it returns a 401 Unauthorized error
		r.Use(CheckAPIKey(svc.Db, svc.Stat))

		// This get request is just a simple ping endpoint to test that the server is running
		// and that the API key is valid.
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		r.Mount("/locations", locations.NewLocationsRouter(&locations.LocationsRouter{Svcs: svc}).Router)
	})

	r.Route("/saml", func(r chi.Router) {
		r.Get("/metadata", samlMiddleware.ServeMetadata)
		r.Post("/acs", samlMiddleware.ServeACS)
	})

	r.Group(func(r chi.Router) {
		// this set of groups require a JWT to be accessed
		r.Use(CheckIdentity(svc.Token, svc.Log))
		r.Mount("/applications", applications.NewApplicationsRouter(&applications.ApplicationsRouter{Svcs: svc}).Router)

	})

	r.Group(func(r chi.Router) {
		r.Use(samlMiddleware.RequireAccount)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			session := samlsp.SessionFromContext(r.Context())
			if session == nil {
				svc.Log.Error(fmt.Sprintf("The context does not contain a session.\n%v", session), nil)
				return
			}

			elon_uid := samlsp.AttributeFromContext(r.Context(), "employeeNumber")
			if elon_uid == "" {
				svc.Log.Error("Elon employeeNumber not provided in context payload.", nil)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if !svc.Db.IsUser(elon_uid) {
				givenName := samlsp.AttributeFromContext(r.Context(), "givenName")
				surname := samlsp.AttributeFromContext(r.Context(), "sn")
				email := samlsp.AttributeFromContext(r.Context(), "mail")
				affiliation := samlsp.AttributeFromContext(r.Context(), "eduPersonPrimaryAffiliation")

				if givenName == "" || surname == "" || email == "" || affiliation == "" {
					svc.Log.Error(fmt.Sprintf("Something's missing... gn: %s, sn: %s, mail: %s, affi: %s", givenName, surname, email, affiliation), nil)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				err := svc.Db.NewUser(elon_uid, givenName, surname, email, affiliation)
				if err != nil {
					svc.Log.Error(err.Error(), nil)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			userInfo, err := svc.Db.GetUserInformation(elon_uid)
			if err != nil {
				svc.Log.Error(err.Error(), nil)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			jwt, err := svc.Token.NewToken(userInfo.OdsId)
			if err != nil {
				svc.Log.Error(err.Error(), nil)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			cleanURL, err := getDomainFromURI(webURL)
			if err != nil {
				svc.Log.Error(err.Error(), nil)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:   "ods_login_cookie_nomnom",
				Value:  jwt,
				MaxAge: 60 * 60 * 2,
				Domain: cleanURL,
			})

			w.Header().Add("Content-Type", "")

			http.Redirect(w, r, webURL+"/apps", http.StatusFound)

		})
	})

	svc.Log.Info("Server running on port "+servicePort, nil)
	svc.Stat.TimeElapsed("server.start", time.Since(startInitialization).Milliseconds())
	return r
}

func main() {
	// get our pertinent information from the environment variables or the command line
	servicePort := flag.String("port", os.Getenv("PORT"), "port to run server on")
	databaseURL := flag.String("database_url", os.Getenv("DATABASE_URL"), "database url")
	redisURL := flag.String("redis_url", os.Getenv("REDIS_URL"), "redis url")
	loggingURL := flag.String("logging_url", os.Getenv("LOGGING_URL"), "logging url")
	statsdURL := flag.String("statsd_url", os.Getenv("STATSD_URL"), "statsd url")
	samlCertPath := flag.String("saml_cert_path", os.Getenv("SAML_CERT_PATH"), "location of service cert")
	samlKeyPath := flag.String("saml_key_path", os.Getenv("SAML_KEY_PATH"), "location of service key")
	idpURL := flag.String("idp_url", os.Getenv("IDP_URL"), "url of identity provider")
	spURL := flag.String("sp_url", os.Getenv("SP_URL"), "url of the hosted service provider")
	webURL := flag.String("web_url", os.Getenv("WEB_URL"), "url of the hosted web service")

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
	if *samlCertPath == "" {
		log.Fatal("service cert location not set")
	}
	if *samlKeyPath == "" {
		log.Fatal("service key location not set")
	}
	if *idpURL == "" {
		log.Fatal("idp url not set")
	}
	if *spURL == "" {
		log.Fatal("sp url not set")
	}

	if *webURL == "" {
		webURL = spURL
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", *servicePort),
		initialize(*servicePort, *databaseURL, *redisURL, *loggingURL, *statsdURL, *samlCertPath, *samlKeyPath, *idpURL, *spURL, *webURL))
	if err != nil {
		fmt.Println(err)
	}
}
