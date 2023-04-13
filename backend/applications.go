package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/elonsoc/center/backend/service"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

/*
	This file deals with application registration and everything that comes with it.
	- recieving the registration form
	- generating a project ID
	- generating an API key
	- storing the information in the database
	- returning the pertinent information to the user

	Database Design can be found: https://docs.google.com/document/d/1zEK9K7crTcCcE9bMKm87qRHCyqa5I0_vjuI9eqYSaLg/edit?usp=sharing
*/

// Defining the struct for the applications router
type ApplicationsRouter struct {
	chi.Router
	Svcs *service.Service
}

// NewApplicationsRouter is a function that returns a new applications router
func NewApplicationsRouter(a *ApplicationsRouter) *ApplicationsRouter {
	a.Svcs.Logger.Info("Initializing applications router")

	r := chi.NewRouter()
	r.Post("/", a.newApp)

	a.Router = r
	a.Svcs.Logger.Info("Applications router initialized")
	return a
}

// create a registration structure (This is not entirely necessary, but it makes things easier)
// Contents of this might change as the project progresses
type application struct {
	projName    string
	projID      string
	description string
	owner       string
	members     string
	apiKey      string
	isValid     bool
}

// This function handles the registration form. It is called when the user submits a registration form.
// It will parse the form, generate a project ID and API key, and store the information in the database.
// It will then return the pertinent information to the user.
func (ar *ApplicationsRouter) newApp(w http.ResponseWriter, r *http.Request) {
	app := application{}
	// parse the request body into the registration variable
	app.apiKey = ar.apiKeyGenerate()
	app.projName = r.FormValue("title")
	app.description = r.FormValue("description")
	app.owner = r.FormValue("owners")
	app.projID = ar.appIDGenerate()
	app.isValid = true
	//app.members = r.FormValue("members") Members are not implemented yet

	// Storing the information in the database. Using a transaction to ensure that the information is stored correctly.
	ctx := context.Background()
	tx, err := ar.Svcs.Db.Begin(ctx)
	if err != nil {
		ar.Svcs.Logger.Info(err)
		ar.Svcs.Logger.Info("Error initiating database transaction")
	}
	defer tx.Rollback(ctx)

	// Storing info in the keys table
	query1 := `INSERT INTO keys (id, apiKey, isValid) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, query1, app.projID, app.apiKey, app.isValid)
	if err != nil {
		ar.Svcs.Logger.Info(err)
		ar.Svcs.Logger.Info("Error storing new app info in database (keys table)")
	}

	// Storing info in the applications table
	query2 := `INSERT INTO applications (id, projName, description, owner, members) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(ctx, query2, app.projID, app.projName, app.description, app.owner, app.members)
	if err != nil {
		ar.Svcs.Logger.Info(err)
		ar.Svcs.Logger.Info("Error storing new app info in database (applications table)")
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		ar.Svcs.Logger.Info(err)
		ar.Svcs.Logger.Info("Error committing database transaction")
	}

}

// TODO: Prepare the response and send it back to the client.
//		 Properly handle errors.

// This function creates a new API key. It is a struct method because it needs to access the database
// and logger. The function will keep generating a new key until it finds one that is unique.
func (ar *ApplicationsRouter) apiKeyGenerate() string {
	// Keep generating a new key until a unique one is found
	isUnique := false
	ctx := context.Background()
	key := ""

	for !isUnique {
		// generate 32 random bytes using crypto/rand
		bytes := make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			ar.Svcs.Logger.Info(err)
			ar.Svcs.Logger.Info("Error generating API key")
		}

		// convert bytes to a base62 string
		base64 := base64.StdEncoding.EncodeToString(bytes)
		base63 := strings.Replace(base64, "+", "0", -1)
		base62 := strings.Replace(base63, "/", "1", -1)

		key = "ods_key_" + base62

		// query the database to see if the key is unique
		query := `SELECT apiKey FROM keys WHERE id=$1`
		rows, err := ar.Svcs.Db.Query(ctx, query, key)
		if err != nil {
			ar.Svcs.Logger.Info(err)
			ar.Svcs.Logger.Info("Error querying database for API key")
		}

		// if the key is unique, set isUnique to true
		if rows.Next() {
			isUnique = false
		} else {
			isUnique = true
		}

		// close the rows
		rows.Close()
	}
	// Return the API key string
	return key
}

// This function creates a new Application ID. It is a struct method because it needs to access the database
// and logger. The function will keep generating a new ID until it finds one that is unique.
func (ar *ApplicationsRouter) appIDGenerate() string {
	// Keep generating a new ID until a unique one is found
	isUnique := false
	appID := ""
	ctx := context.Background()

	for !isUnique {
		// generate 32 random bytes using crypto/rand
		bytes := make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			ar.Svcs.Logger.Info(err)
			ar.Svcs.Logger.Info("Error generating App ID")
		}

		// convert bytes to a base62 string
		base64 := base64.StdEncoding.EncodeToString(bytes)
		base63 := strings.Replace(base64, "+", "0", -1)
		base62 := strings.Replace(base63, "/", "1", -1)

		appID = "ods_app_" + base62

		// query the database to see if the key is unique
		query := `SELECT id FROM keys WHERE id=$1`
		rows, err := ar.Svcs.Db.Query(ctx, query, appID)
		if err != nil {
			ar.Svcs.Logger.Info(err)
			ar.Svcs.Logger.Info("Error querying database for App ID")
		}

		// if the appID is unique, set isUnique to true
		if rows.Next() {
			isUnique = false
		} else {
			isUnique = true
		}

		// close the rows
		rows.Close()
	}
	// Return the API key string
	return appID
}

/*
All the code below has not been implemented yet or extensively worked on.

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

*/
