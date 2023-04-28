package applications

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
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
	a.Svcs.Log.Info("Initializing applications router", nil)

	r := chi.NewRouter()
	r.Post("/", a.newApp)
	r.Get("/", a.myApps)

	a.Router = r
	a.Svcs.Log.Info("Applications router initialized", nil)
	return a
}

// The Application type defines the structure of an application.
type Application struct {
	AppName     string `json:"appName" db:"app_name"`
	AppID       string `json:"appID" db:"app_ID"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners" db:"owners"`
	TeamName    string `json:"teamName" db:"team_name"`
	ApiKey      string `json:"apiKey" db:"api_key"`
	IsValid     bool   `json:"isValid" db:"is_valid"`
}

// This function handles the registration form. It is called when the user submits a registration form.
// It will parse the form, generate a project ID and API key, and store the information in the database.
// It will then return the pertinent information to the user.
func (ar *ApplicationsRouter) newApp(w http.ResponseWriter, r *http.Request) {
	var err error
	// Create a new Application struct
	app := Application{}
	// parse the request body into the application variable
	app.AppName = r.FormValue("title")
	app.Description = r.FormValue("description")
	app.Owners = r.FormValue("owners")
	app.TeamName = r.FormValue("teamName")
	app.IsValid = true
	// Generate a new AppID and API key
	app.AppID, err = ar.appIDGenerate()
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}
	app.ApiKey, err = ar.apiKeyGenerate()
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	// Store the application in the database
	err = ar.Svcs.Db.NewApp(app.AppName, app.AppID, app.Description, app.Owners, app.TeamName, app.ApiKey, app.IsValid)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}
}

/* THIS IS A DUMMY END POINT RIGHT NOW */
// This function returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return a
// list of all the applications that the user owns.
func (ar *ApplicationsRouter) myApps(w http.ResponseWriter, r *http.Request) {
	// Query db for all apps
	rows, err := ar.Svcs.Db.UserApps()
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	// Scan the rows into an array of Applications
	apps := []Application{}
	for rows.Next() {
		app := Application{}
		err = rows.Scan(&app.AppName, &app.AppID, &app.Description, &app.Owners, &app.TeamName)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return
		}
		apps = append(apps, app)
	}

	// Encode the apps array into JSON and send it back to the client
	json.NewEncoder(w).Encode(apps)
}

// This function creates a new API key. It is a struct method because it needs to access the database
// and logger. The function will keep generating a new key until it finds one that is unique.
func (ar *ApplicationsRouter) apiKeyGenerate() (string, error) {
	// Keep generating a new key until a unique one is found
	isUnique := false
	key := ""

	for !isUnique {
		// generate 32 random bytes using crypto/rand
		bytes := make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return "", err
		}

		// convert bytes to a base62 string
		base64 := base64.StdEncoding.EncodeToString(bytes)
		base63 := strings.Replace(base64, "+", "0", -1)
		base62 := strings.Replace(base63, "/", "1", -1)

		key = "ods_key_" + base62

		// query the database to see if the key is unique
		isUnique, err = ar.Svcs.Db.CheckDuplicate("api_key", key)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return "", err
		}
	}
	// Return the API key string
	return key, nil
}

// This function creates a new Application ID. It is a struct method because it needs to access the database
// and logger. The function will keep generating a new ID until it finds one that is unique.
func (ar *ApplicationsRouter) appIDGenerate() (string, error) {
	// Keep generating a new ID until a unique one is found
	isUnique := false
	appID := ""

	for !isUnique {
		// generate 32 random bytes using crypto/rand
		bytes := make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return "", err
		}

		// convert bytes to a base62 string
		base64 := base64.StdEncoding.EncodeToString(bytes)
		base63 := strings.Replace(base64, "+", "0", -1)
		base62 := strings.Replace(base63, "/", "1", -1)

		appID = "ods_app_" + base62

		// query the database to see if the key is unique
		isUnique, err = ar.Svcs.Db.CheckDuplicate("app_ID", appID)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return "", err
		}
	}
	// Return the API key string
	return appID, nil
}

/*
All the code below has not been implemented yet or extensively worked on.

func apiKeyRefresh(appID string) {
	// generate a new key with apiKeyGenerate()
	sqlQuery := "UPDATE Keys SET \"Api Key\" = $1 WHERE \"Project ID\" = $2"
	// store new key in database
	res, err := db.Exec(sqlQuery, apiKeyGenerate(), appID)
	if err != nil {
		// log error
	}
}

func apiKeyRevoke() {
	// update "isValid" in the database to false
}

*/
