package apps

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"net/http"
)

// create a registration structure
// more can be added to this later (api key for example)
type project struct {
	projName    string
	projID      string
	description string
	owner       string
	members     []string
	apiKey      string
}

// Everytime a registration is filled out, a new registration variable is created
// This will then be passed through to the DB once the DB is figured out

// Database Design can be found: https://docs.google.com/document/d/1zEK9K7crTcCcE9bMKm87qRHCyqa5I0_vjuI9eqYSaLg/edit?usp=sharing

func newApp(w http.ResponseWriter, r *http.Request) {

}

func apiKeyGenerate() string {
	// generate 36 random bytes using crypto/rand
	bytes := make([]byte, 36)
	_, err := rand.Read(bytes)
	if err != nil {
		// log error
	}

	// convert bytes to a base62 string
	base64 := base64.StdEncoding.EncodeToString(bytes)
	base63 := strings.Replace(base64, "+", "0", -1)
	base62 := strings.Replace(base63, "/", "1", -1)

	// Return the API key string
	return "ods_App_" + base62
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
