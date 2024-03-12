package token

import (
	"encoding/json"
	"net/http"

	"github.com/elonsoc/ods/src/auth/pkg"
	"github.com/elonsoc/ods/src/common"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func SetTokenEndpoint(svc *common.Services, tok TokenIFace) chi.Router {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		t := struct{ Uid string }{}
		json.NewDecoder(r.Body).Decode(&t)
		newToken, err := tok.NewToken(t.Uid)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res, err := json.Marshal(pkg.TokenHttp{T: newToken})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(res)
	})

	r.Post("/validate", func(w http.ResponseWriter, r *http.Request) {
		t := pkg.TokenHttp{}
		json.NewDecoder(r.Body).Decode(&t)
		verdict, err := tok.ValidateToken(t.T)
		if err != nil {
			svc.Log.Error(err.Error(), logrus.Fields{"token": t.T})
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res, err := json.Marshal(struct {
			Verdict bool `json:"verdict"`
		}{Verdict: verdict})
		if err != nil {
			svc.Log.Error(err.Error(), nil)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(res)

	})

	r.Post("/uid", func(w http.ResponseWriter, r *http.Request) {
		t := pkg.TokenHttp{}
		json.NewDecoder(r.Body).Decode(&t)
		uid, err := tok.GetUidFromToken(t.T)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(struct{ Uid string }{uid})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(res)
	})

	r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
		t := pkg.TokenHttp{}
		json.NewDecoder(r.Body).Decode(&t)
		accessToken, refreshToken, err := tok.RefreshAccessToken(t.T)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(struct {
			Access  string
			Refresh string
		}{Access: accessToken, Refresh: refreshToken})
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(res)
	})

	r.Post("/invalidate", func(w http.ResponseWriter, r *http.Request) {
		t := pkg.TokenHttp{}
		json.NewDecoder(r.Body).Decode(&t)
		err := tok.InvalidateToken(t.T)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	return r
}
