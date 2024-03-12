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

import "github.com/elonsoc/ods/src/common"

// Service, here, describes the services that we will be using
// in the backend of ods.
type Services struct {
	Log   common.LoggerIFace
	Db    DbIFace
	Stat  StatIFace
	Saml  SamlIFace
	Token TokenIFace
	Redis InMemoryDbIFace
}

// NewService creates a new instance of the Service struct
// and returns a pointer to it.
// This is necessary because we want to be able to pass the
// Service struct around by reference, and not by value.
// If we pass it by value, we would be passing a copy of the
// struct, and any changes made to the struct would not be
// reflected in the original struct and thus not able to
// be used by other functions.
func NewService(loggingURL, databaseURL, redisURL, statsdURL, certPath, keyPath, idpURL, spURL, webURL string) *Services {
	// We are using the log package here to create a new logger
	// that will be used to log messages to the console.

	log := common.InitLogging(loggingURL)
	db := initDb(databaseURL, log)
	stat := initStatsD(statsdURL, log)
	saml := initializeSaml(log, idpURL, webURL, spURL, certPath, keyPath)
	redis := initInMemoryDb(redisURL, log)
	token := NewTokenServicer(redis)

	return &Services{
		Log:   log,
		Db:    db,
		Stat:  stat,
		Saml:  saml,
		Token: token,
		Redis: redis,
	}
}
