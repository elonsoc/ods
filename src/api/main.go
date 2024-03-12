package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	statsd "github.com/smira/go-statsd"

	"github.com/elonsoc/ods/src/api/applications"
	locations "github.com/elonsoc/ods/src/api/locations"
	auth "github.com/elonsoc/ods/src/auth/pkg"
	"github.com/elonsoc/ods/src/common"
)

// CheckAuth is a custom middleware that checks the Authorization header for a valid API key
func CheckAPIKey(db common.DbIFace, stat common.StatIFace) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				http.Error(w, "Authorization header must be in the format 'Bearer {token}'", http.StatusUnauthorized)
				return
			}

			apiKey := parts[1]

			// this is where the call we'd make to the database to verify the API key
			// would happen. For now, we just check if the API key is in the map above.
			if !db.IsValidApiKey(apiKey) {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}

			// if both are okay, then we're free to further this request.
			stat.Increment("api_key_used", statsd.StringTag("api_key", apiKey))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func CheckIdentity(tokenSvcr auth.TokenIFace, log common.LoggerIFace) func(next http.Handler) http.Handler {
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

// initialize begins the startup process for the backend of ods.
// At the beginning, we create a new instance of the router, declare usage of multiple middlewares
// initialize connections to external services, and mount the various routers for the apis that we
// will be serving.
func initialize(urls common.Flags) chi.Router {
	startInitialization := time.Now()
	// This is where we initialize the various services that we will be using
	// like the database, logger, stats, etc.
	// in the future, we could split up the apis into separate services
	// which would allow for more flexibility in scaling, allow for more granular control over the API keys,
	// and allow for more granular control over the services that are running.
	// This particular technique is called dependency injection, and it's a good practice to use
	// when writing code that could one day be decoupled into separate services.
	// There are better ways to do this, but this is a good start to keep the app monolithic for now.
	svc := common.NewService(urls)
	tokenSvc := auth.NewTokenService(*urls.AuthURL)

	// Create a new instance of the router
	r := chi.NewRouter()

	// Middleware is defined next and works in a LIFO (Last In First Out) order.
	// This means that the first middleware that is declared will be the last one to be executed.

	// middleware.Logger prints a log line for each request (access log)
	r.Use(common.CustomLogger(svc.Log, svc.Stat))
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

	r.Group(func(r chi.Router) {
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			for _, cookie := range r.Cookies() {
				switch cookie.Name {
				case "ods_login_cookie_nomnom", "ods_refresh_cookie_nomnom":
					tokenPrefix := map[string]string{
						"ods_login_cookie_nomnom":   "access_token:",
						"ods_refresh_cookie_nomnom": "refresh_token:",
					}[cookie.Name]

					odsId, err := tokenSvc.GetUidFromToken(cookie.Value)
					if err != nil {
						svc.Log.Error("Failed to get odsId from token: "+err.Error(), nil)
						http.Error(w, "Server error", http.StatusInternalServerError)
						return
					}
					tokenKey := tokenPrefix + odsId
					if err = tokenSvc.InvalidateToken(tokenKey); err != nil {
						svc.Log.Error("Failed to invalidate token: "+tokenKey+" "+err.Error(), nil)
						http.Error(w, "Server error", http.StatusInternalServerError)
						return
					}
					svc.Log.Info("Token invalidated: "+tokenKey, nil)
				}

				cookie.MaxAge = -1
				http.SetCookie(w, cookie)
			}
			http.Redirect(w, r, *urls.ServicePort, http.StatusFound)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(CheckIdentity(tokenSvc, svc.Log))
		r.Mount("/applications", applications.NewApplicationsRouter(&applications.ApplicationsRouter{Svcs: svc}).Router)
		r.Get("/login/status", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
		refreshToken := r.Header.Get("X-Refresh-Token")
		svc.Log.Info("refresh token: "+refreshToken, nil)
		if refreshToken == "" {
			svc.Log.Error("No refresh token provided", nil)
			http.Error(w, "No refresh token provided", http.StatusBadRequest)
			return
		}

		newAccessToken, newRefreshToken, err := tokenSvc.RefreshAccessToken(refreshToken)
		if err != nil {
			svc.Log.Error("Failed to refresh access token: "+err.Error(), nil)
			http.Error(w, "Failed to refresh access token", http.StatusInternalServerError)
			return
		}

		tokens := struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token,omitempty"`
		}{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokens)
	})

	svc.Log.Info("Server running on port "+*urls.ServicePort, nil)
	svc.Stat.TimeElapsed("server.start", time.Since(startInitialization).Milliseconds())
	return r
}

func main() {
	urls := common.GetAndParseFlags()

	err := http.ListenAndServe(fmt.Sprintf(":%s", *urls.ServicePort),
		initialize(urls))
	if err != nil {
		fmt.Println(err)
	}
}
