package service

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

// DbIFace is an interface for the database
type DbIFace interface {
	GetConn() *pgx.Conn
	NewApp(string, string, string, string, string, string, bool) error
	UserApps() (pgx.Rows, error)
	CheckDuplicate(string, string) (bool, error)
	GetApplication(string) (ApplicationExtended, error)
	UpdateApplication(string, ApplicationSimple) error
	DeleteApplication(string) error
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgx.Conn
}

type ApplicationSimple struct {
	Id          string `json:"appID" db:"app_ID"`
	Name        string `json:"appName" db:"app_name"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners" db:"owners"`
	Team        string `json:"teamName" db:"team_name"`
}

type ApplicationExtended struct {
	Id          string `json:"appID" db:"app_ID"`
	Name        string `json:"appName" db:"app_name"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners" db:"owners"`
	Team        string `json:"teamName" db:"team_name"`
	Api_key     string `json:"apiKey" db:"api_key"`
	Is_valid    bool   `json:"isValid" db:"is_valid"`
}

func initDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()
	var err error
	conn, err = pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	err = prepareStatements(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func prepareStatements(connection *pgx.Conn, ctx context.Context) (err error) {
	_, err = connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (app_ID, app_name, description, owners, team_name, api_key, is_valid) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	return nil
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

// NewApp stores the information about a new application in the database.
func (db *Db) NewApp(name string, ID string, desc string, owners string, tname string, key string, valid bool) error {
	ctx := context.Background()

	// Storing all new app info into the applications table.
	db.db.Exec(ctx, "insert_into_applications", ID, name, desc, owners, tname, key, valid)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) UserApps() (pgx.Rows, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "SELECT * FROM applications")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *Db) GetApplication(applicationId string) (ApplicationExtended, error) {
	row, _ := conn.Query(context.Background(), "SELECT * FROM applications WHERE app_ID=$1", applicationId)

	var app ApplicationExtended

	for row.Next() {
		err := row.Scan(&app.Id,
			&app.Name,
			&app.Description,
			&app.Owners,
			&app.Team,
			&app.Api_key,
			&app.Is_valid)
		if err != nil {
			return app, err
		}
	}

	return app, nil
}

func (db *Db) UpdateApplication(applicationId string, applicationInfo ApplicationSimple) error {
	_, err := conn.Exec(context.Background(), "UPDATE applications SET app_name=$1, description=$2, owners=$3, team_name=$4 WHERE app_ID=$5",
		applicationInfo.Name,
		applicationInfo.Description,
		applicationInfo.Owners,
		applicationInfo.Team,
		applicationId)
	return err
}

func (db *Db) DeleteApplication(applicationId string) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM applications WHERE app_ID=$1", applicationId)
	return err
}

// Checks for duplicated app IDs or api keys
// I easily could've used if statements but I've always wanted to use switch statements
// and realistically this could expand to more columns in the future.
func (db *Db) CheckDuplicate(column string, newGen string) (bool, error) {
	ctx := context.Background()
	var query string

	// This switch statement is used to determine which column to check for duplicates
	switch column {
	case "app_ID":
		query = "SELECT * FROM applications WHERE app_ID = $1"
	case "api_key":
		query = "SELECT * FROM applications WHERE api_key = $1"
	}

	// Querying the database for duplicates
	err := db.db.QueryRow(ctx, query, newGen).Scan()
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
