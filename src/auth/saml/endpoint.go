package saml

import (
	"fmt"
	"net/http"

	"github.com/elonsoc/ods/src/auth/token"
	"github.com/elonsoc/ods/src/common"
	"github.com/elonsoc/saml/samlsp"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetSamlEndpoint(webURL string, svc *common.Services, t token.TokenIFace, smw *samlsp.Middleware) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Get("/metadata", smw.ServeMetadata)
		r.Post("/acs", smw.ServeACS)
	})

	r.Group(func(r chi.Router) { // prolly fwd all of this to auth service
		r.Use(smw.RequireAccount)
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

			svc.Log.Info(fmt.Sprintf("we're given a elon_uid: %s", elon_uid), nil)

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

			jwt, err := t.GenerateAccessToken(userInfo.OdsId)
			if err != nil {
				svc.Log.Error(err.Error(), nil)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			refreshToken, err := t.GenerateRefreshToken(userInfo.OdsId)
			if err != nil {
				svc.Log.Error("Failed to generate refresh token: "+err.Error(), nil)
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
				Name:     "ods_login_cookie_nomnom",
				Value:    jwt,
				MaxAge:   60 * 5, /* * 60 * 2 */
				Domain:   cleanURL,
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
			})

			http.SetCookie(w, &http.Cookie{
				Name:     "ods_refresh_cookie_nomnom",
				Value:    refreshToken,
				MaxAge:   60 * 60 * 24 * 7,
				Path:     "/",
				Domain:   cleanURL,
				Secure:   true,
				HttpOnly: true,
			})

			w.Header().Add("Content-Type", "")

			http.Redirect(w, r, webURL+"/apps", http.StatusFound)

		})
	})
	return r
}
