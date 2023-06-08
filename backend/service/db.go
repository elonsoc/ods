package service

import (
	"context"

	"github.com/elonsoc/ods/backend/applications/types"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DbIFace is an interface for the database
type DbIFace interface {
	NewApp(string, string) (string, error)
	GetApplications() ([]types.Application, error)
	GetApplication(string) (types.Application, error)
	UpdateApplication(string, types.Application) error
	DeleteApplication(string) error
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgxpool.Pool
}

const (
	insertIntoApplications     = "INSERT INTO applications (name, description, is_valid) VALUES ($1, $2, true) RETURNING id"
	selectAllApplications      = "SELECT api_key, description, id, is_valid, name FROM applications"
	selectApplication          = "SELECT api_key, description, id, is_valid, name FROM applications WHERE id=$1"
	getOwners                  = "SELECT email FROM users WHERE id IN (SELECT user_id FROM app_owner WHERE app_id=$1)"
	updateApplication          = "UPDATE applications SET name=$1, description=$2 WHERE id=$3"
	firstLegDeleteApplication  = "DELETE FROM app_owner WHERE app_id=$1"
	secondLegDeleteApplication = "DELETE FROM applications WHERE id=$1"
)

func initDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: pool}
}

// NewApp stores the information about a new application in the database.

func (db *Db) NewApp(name string, desc string) (string, error) {
	ctx := context.Background()

	var id string

	err := db.db.QueryRow(ctx, insertIntoApplications, name, desc).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) GetApplications() ([]types.Application, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, selectAllApplications)
	if err != nil {
		return nil, err
	}
	apps := []types.Application{}
	for rows.Next() {
		var app types.Application
		err = rows.Scan(&app.ApiKey, &app.Description, &app.Id, &app.IsValid, &app.Name)
		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}
	return apps, nil
}

func (db *Db) GetApplication(applicationId string) (types.Application, error) {
	var app types.Application
	err := db.db.QueryRow(context.Background(), selectApplication, applicationId).Scan(
		&app.ApiKey, &app.Description, &app.Id, &app.IsValid, &app.Name)
	if err != nil {
		return app, err
	}

	rows, err := db.db.Query(context.Background(), getOwners, applicationId)
	if err != nil {
		return app, err
	}

	if rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return app, err
		}
		app.Owners = email
	}

	return app, nil
}

func (db *Db) UpdateApplication(applicationId string, applicationInfo types.Application) error {
	_, err := db.db.Exec(context.Background(), updateApplication,
		applicationInfo.Name,
		applicationInfo.Description,
		applicationId)
	return err
}

func (db *Db) DeleteApplication(applicationId string) error {
	txContext := context.Background()
	tx, err := db.db.BeginTx(txContext, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(txContext)

	_, err = tx.Exec(txContext, firstLegDeleteApplication, applicationId)
	if err != nil {
		return err
	}
	tx.Exec(txContext, secondLegDeleteApplication, applicationId)
	if err != nil {
		return err
	}
	err = tx.Commit(txContext)
	if err == nil {
		return nil
	}
	return err
}
