package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	locations "github.com/elonsoc/center/locations"
	"github.com/elonsoc/center/service"
)

func CheckAuth(next http.Handler) http.Handler {
	// this mocks a database of API keys
	var APIKEYS = map[string]bool{
		"12345": true,
	}

	// this middleware is a proof of concept, in the future we would swap this out
	// for a call to our database to verify the API key
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking auth")
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

		next.ServeHTTP(w, r)
	})
}

// initialize begins the startup process for the backend of launchpad.
// At the beginning, we create a new instance of the router, declare usage of multiple middlewares
// initialize connections to external services, and mount the various routers for the apis that we
// will be serving.
func initialize() chi.Router {
	// Create a new instance of the router
	r := chi.NewRouter()

	// Middleware is defined next and works in a LIFO (Last In First Out) order.
	// This means that the first middleware that is declared will be the last one to be executed.

	// middleware.Logger prints a log line for each request (access log)
	r.Use(middleware.Logger)
	// middleware.Recoverer recovers from panics without crashing the server and supplies
	// a 500 error page (Internal Server Error)
	r.Use(middleware.Recoverer)
	// middleware.RealIP is used to get the real IP address of the client
	r.Use(middleware.RealIP)
	// This custom middleware checks the Authorization header for a valid API key
	// and if it's not valid, it returns a 401 Unauthorized error
	r.Use(CheckAuth)

	// This GET request covers the index of the API, just to be a bit cute.
	// Though the API is not meant to be accessed by a browser, this is a good
	// way of verifying that your API key is working.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("If you're seeing this, you're authenticated!"))
	})

	// This is where we initialize the various services that we will be using
	// like the database, logger, stats, etc.
	// in the future, we could split up the apis into separate services
	// which would allow for more flexibility in scaling, allow for more granular control over the API keys,
	// and allow for more granular control over the services that are running.
	// This particular technique is called dependency injection, and it's a good practice to use
	// when writing code that could one day be decoupled into separate services.
	// There are better ways to do this, but this is a good start to keep the app monolithic for now.
	Services := service.NewService()

	// This is where we mount the various routers for the various APIs that we will be serving
	// again, in the future we can split these up into separate services if we want to for the above reasons
	// routers are versioned, so we can have multiple versions of the same API running at the same time
	// which is a good practice for API design to allow for backwards compatibility and support for
	// older clients who may rely on older versions of the API.
	r.Mount("/locations", locations.NewLocationsRouter(&locations.LocationsRouter{Svcs: Services}).Router)

	return r
}

func main() {
	// This is where we start the server and listen on port 1337
	err := http.ListenAndServe(":1337", initialize())
	if err != nil {
		fmt.Println(err)
	}
}
