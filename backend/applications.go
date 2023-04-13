package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/elonsoc/center/backend/service"
	"github.com/go-chi/chi/v5"
)

// Defining the struct for the applications router
type ApplicationsRouter struct {
	chi.Router
	Svcs *service.Service
}

// NewApplicationsRouter is a function that returns a new applications router
func NewApplicationsRouter(a *ApplicationsRouter) *ApplicationsRouter {
	a.Svcs.Logger.Info("Initializing applications router")

	r := chi.NewRouter()
	r.Post("/", newApp)

	a.Router = r
	a.Svcs.Logger.Info("Applications router initialized")
	return a
}

// create a registration structure
// more can be added to this later (api key for example)
type application struct {
	projName    string
	projID      string
	description string
	owner       string
	members     []string
	apiKey      string
}

// Everytime a registration is filled out, a new application variable is created
// This will then be passed through to the DB once the DB is figured out

// Database Design can be found: https://docs.google.com/document/d/1zEK9K7crTcCcE9bMKm87qRHCyqa5I0_vjuI9eqYSaLg/edit?usp=sharing

func (ar *ApplicationsRouter) newApp(w http.ResponseWriter, r *http.Request) {
	app := application{}
	// parse the request body into the registration variable
	app.apiKey = apiKeyGenerate()
	app.projName = r.FormValue("title")
	app.description = r.FormValue("description")
	app.owner = r.FormValue("owners")
	app.projID = 

	// TODO: Create Project ID function.
	//       Make apiKey and projID generation functions struct methods.
	//       This means we can use the database (and logger) connection in the struct.
	//       This allows us to check if the generated string is unique.
	//       Once that is done, we can store the new application info in the database.
	//       Prepare the response and send it back to the client.
	
	
}

func apiKeyGenerate() string {
	// generate 32 random bytes using crypto/rand
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		// log error
	}

	// convert bytes to a base62 string
	base64 := base64.StdEncoding.EncodeToString(bytes)
	base63 := strings.Replace(base64, "+", "0", -1)
	base62 := strings.Replace(base63, "/", "1", -1)

	// Return the API key string
	return "ods_key_" + base62
}

func apiKeyRefresh(projID string) {
	// generate a new key with apiKeyGenerate()
	sqlQuery := "UPDATE Keys SET \"Api Key\" = $1 WHERE \"Project ID\" = $2"
	// store new key in database
	res, err := db.Exec(sqlQuery, apiKeyGenerate(), projID)
	if err != nil {
		// log error
	}
}

func apiKeyRevoke() {
	// update "isValid" in the database to false
}
