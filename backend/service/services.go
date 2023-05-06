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
	"database/sql"
	"github.com/jackc/pgx/v5"
	_ "github.com/microsoft/go-mssqldb"
	logrusLoki "github.com/schoentoon/logrus-loki"
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"
)

// Service, here, describes the services that we will be using
// in the backend of ods.
type Service struct {
	Logger *logrus.Logger
	PgDb   *pgx.Conn
	MsDb   *sql.DB
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
func NewService(loggingURL, pgURL, mssqlURL, statsdURL string) *Service {
	// We are using the log package here to create a new logger
	// that will be used to log messages to the console.

	log := logrus.New()
	hook, err := logrusLoki.NewLokiDefaults(loggingURL)
	if err != nil {
		log.Fatal(err)
	}
	log.AddHook(hook)

	pgdb := InitPgDB(pgURL, log)
	msdb := InitMsDB(mssqlURL, log)

	stat := InitStatsD(statsdURL, log)

	return &Service{
		Logger: log,
		PgDb:   pgdb,
		MsDb:   msdb,
		Stat:   stat,
	}
}

func InitStatsD(statsdURL string, log *logrus.Logger) *statsd.Client {
	client := statsd.NewClient(statsdURL, statsd.MetricPrefix("backend."))
	return client
}

func InitPgDB(pgdbURL string, log *logrus.Logger) *pgx.Conn {
	connection, err := pgx.Connect(context.Background(), pgdbURL)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func InitMsDB(msdbURL string, log *logrus.Logger) *sql.DB {
	db, err := sql.Open("sqlserver", msdbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
