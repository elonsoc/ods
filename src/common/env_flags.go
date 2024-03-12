package common

import (
	"flag"
	"log"
	"os"
)

type Flags struct {
	ServicePort *string
	DbURL       *string
	IMDbURL     *string
	LogURL      *string
	StatURL     *string
	WebURL      *string
	AuthURL     *string
}

// GetAndParseFlags parses the flags common amongst all backend services
// you must call this after setting the service's scoped flags.
func GetAndParseFlags() Flags {
	// get our pertinent information from the environment variables or the command line
	servicePort := flag.String("port", os.Getenv("PORT"), "port to run server on")
	databaseURL := flag.String("database_url", os.Getenv("DATABASE_URL"), "database url")
	redisURL := flag.String("redis_url", os.Getenv("REDIS_URL"), "redis url")
	loggingURL := flag.String("logging_url", os.Getenv("LOGGING_URL"), "logging url")
	statsdURL := flag.String("statsd_url", os.Getenv("STATSD_URL"), "statsd url")
	webURL := flag.String("web_url", os.Getenv("WEB_URL"), "url of the hosted web service")
	authURL := flag.String("auth_url", os.Getenv("AUTH_URL"), "url of the auth service")

	flag.Parse()
	if *servicePort == "" {
		log.Fatal("port not set")
	}
	if *databaseURL == "" {
		log.Fatal("database url not set")
	}
	if *redisURL == "" {
		log.Fatal("redis url not set")
	}
	if *loggingURL == "" {
		log.Fatal("logging url not set")
	}
	if *statsdURL == "" {
		log.Fatal("statsd url not set")
	}
	if *webURL == "" {
		log.Fatal("web url not set")
	}
	if *authURL == "" {
		log.Fatal("auth url not set.")
	}

	return Flags{
		ServicePort: servicePort,
		DbURL:       databaseURL,
		IMDbURL:     redisURL,
		LogURL:      loggingURL,
		StatURL:     statsdURL,
		WebURL:      webURL,
		AuthURL:     authURL,
	}
}
