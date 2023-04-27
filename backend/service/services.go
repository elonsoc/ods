// This file contains the definition of the Service struct
// which is used to define the various services that we will be using
// in a dependency injection pattern.
// A dependency injection pattern defines the services that we will be using
// in a generalized way, and then we can pass the services around by reference
// to the various functions that need them.
// This allows us to define the services in one place, and then use them
// in any function that needs them as well as change the services that we are using
// in one place, and have those changes reflected in all of the functions that use them.

// By separating the services into their own file, we can focus the different packages
// on their specific tasks, and not have to worry about the services that they are using.

package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	logrusLoki "github.com/schoentoon/logrus-loki"
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"
)

// Service, here, describes the services that we will be using
// in the backend of ods.
type Service struct {
	Logger *logrus.Logger
	Db     *pgx.Conn
	Stat   *statsd.Client
}

// NewService creates a new instance of the Service struct
// and returns a pointer to it.
// This is necessary because we want to be able to pass the
// Service struct around by reference, and not by value.
// If we pass it by value, we would be passing a copy of the
// struct, and any changes made to the struct would not be
// reflected in the original struct and thus not able to
// be used by other functions.
func NewService(loggingURL, databaseURL, statsdURL string) *Service {
	// We are using the log package here to create a new logger
	// that will be used to log messages to the console.

	log := logrus.New()
	hook, err := logrusLoki.NewLokiDefaults(loggingURL)
	if err != nil {
		log.Fatal(err)
	}
	log.AddHook(hook)
	db := InitDB(databaseURL, log)

	stat := InitStatsD(statsdURL, log)

	return &Service{
		Logger: log,
		Db:     db,
		Stat:   stat,
	}
}

func InitStatsD(statsdURL string, log *logrus.Logger) *statsd.Client {
	client := statsd.NewClient(statsdURL, statsd.MetricPrefix("backend."))
	return client
}

func InitDB(databaseURL string, log *logrus.Logger) *pgx.Conn {
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Preparing a few statements that we will be using in other functions and files.
	// Presumably this will be expanded on later and really should be better organized.
	connection.Prepare(ctx, "insert_into_keys", "INSERT INTO keys (app_ID, api_key, is_valid) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (app_ID, app_name, description, owners, team_name) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
