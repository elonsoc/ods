package applications

import (
	"encoding/json"
	"net/http"

	"github.com/elonsoc/ods/backend/applications/types"
	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

/*
	This file deals with all things applications
	- recieving the registration form
	- generating a project ID and an API key
	- storing the information in the database
	- querying the database for the user's applications
	- revoking an API key
	- refreshing an API key

	Database Design can be found: https://docs.google.com/document/d/1zEK9K7crTcCcE9bMKm87qRHCyqa5I0_vjuI9eqYSaLg/edit?usp=sharing
*/

// Defining the struct for the applications router
type ApplicationsRouter struct {
	chi.Router
	Svcs *service.Services
}

// NewApplicationsRouter is a function that returns a new applications router
func NewApplicationsRouter(a *ApplicationsRouter) *ApplicationsRouter {
	r := chi.NewRouter()

	// CORS is used to allow cross-origin requests
	cors := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		MaxAge:             3599, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Post("/", a.newApp)
	r.Get("/", a.myApps)
	r.Route("/{id}", func(r chi.Router) {
		r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		r.Get("/", a.GetApplication)
		r.Put("/", a.UpdateApplication)
		r.Delete("/", a.DeleteApplication)
	})

	a.Router = r
	return a
}

// This function handles the registration form. It is called when the user submits a registration form.
// It will parse the form, generate a project ID and API key, and store the information in the database.
// It will then return the pertinent information to the user.
func (ar *ApplicationsRouter) newApp(w http.ResponseWriter, r *http.Request) {
	var err error

	token, err := r.Cookie("ods_login_cookie_nomnom")
	if err != nil {
		// how did you get here?
		ar.Svcs.Log.Error(err.Error(), nil)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uid, err := ar.Svcs.Token.GetUidFromToken(token.Value)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	// Create a new Application struct
	app := types.BaseApplication{}

	err = json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	// Store the application in the database
	appId, err := ar.Svcs.Db.NewApp(app.Name, app.Description, uid)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(appId))
}

// myApps get all of the apps for a particular user id.
func (ar *ApplicationsRouter) myApps(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("ods_login_cookie_nomnom")
	if err != nil {
		// how did you get here?
		ar.Svcs.Log.Error(err.Error(), nil)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uid, err := ar.Svcs.Token.GetUidFromToken(token.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	if uid == "" {
		w.WriteHeader(http.StatusUnauthorized)
		ar.Svcs.Log.Error("uid is empty", nil)
		return
	}

	apps, err := ar.Svcs.Db.GetApplications(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	json.NewEncoder(w).Encode(apps)
}

func (ar *ApplicationsRouter) GetApplication(w http.ResponseWriter, r *http.Request) {
	applicationId := chi.URLParam(r, "id")

	app, err := ar.Svcs.Db.GetApplication(applicationId)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), logrus.Fields{"applicationId": applicationId})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(app)
}

func (ar *ApplicationsRouter) UpdateApplication(w http.ResponseWriter, r *http.Request) {
	applicationId := chi.URLParam(r, "id")
	var err error

	app := types.Application{}
	app.Id = applicationId
	err = json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	err = ar.Svcs.Db.UpdateApplication(applicationId, app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ar *ApplicationsRouter) DeleteApplication(w http.ResponseWriter, r *http.Request) {
	applicationId := chi.URLParam(r, "id")
	ar.Svcs.Log.Info("DeleteApplication", logrus.Fields{"applicationId": applicationId})
	err := ar.Svcs.Db.DeleteApplication(applicationId)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), logrus.Fields{"applicationId": applicationId})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
