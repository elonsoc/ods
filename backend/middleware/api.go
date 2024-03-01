package middleware

import (
	"net/http"

	"github.com/elonsoc/ods/backend/service"
	"github.com/smira/go-statsd"
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
