package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elonsoc/ods/src/auth/saml"
	"github.com/elonsoc/ods/src/auth/token"
	"github.com/elonsoc/ods/src/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initialize(urls common.Flags, certPath, keyPath, idpURL, spURL string) chi.Router {
	// startInitialization := time.Now()
	svc := common.NewService(urls)
	s := saml.InitializeSaml(svc, idpURL, *urls.WebURL, spURL, certPath, keyPath)

	t := token.NewTokenServicer(svc.IMDb)

	smw := s.Middleware()

	r := chi.NewRouter()
	r.Use(common.CustomLogger(svc.Log, svc.Stat))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Mount("/saml", saml.SetSamlEndpoint(*urls.WebURL, svc, t, smw))
	r.Mount("/token", token.SetTokenEndpoint(svc, t))

	return r
}
func main() {
	samlCertPath := flag.String("saml_cert_path", os.Getenv("SAML_CERT_PATH"), "location of service cert")
	samlKeyPath := flag.String("saml_key_path", os.Getenv("SAML_KEY_PATH"), "location of service key")
	idpURL := flag.String("idp_url", os.Getenv("IDP_URL"), "url of identity provider")
	spURL := flag.String("sp_url", os.Getenv("SP_URL"), "url of the hosted service provider")

	f := common.GetAndParseFlags()
	if *samlCertPath == "" {
		log.Fatal("service cert location not set")
	}
	if *samlKeyPath == "" {
		log.Fatal("service key location not set")
	}
	if *idpURL == "" {
		log.Fatal("idp url not set")
	}
	if *spURL == "" {
		log.Fatal("sp url not set")
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", *f.ServicePort),
		initialize(f, *samlCertPath, *samlKeyPath, *idpURL, *spURL))
	if err != nil {
		fmt.Println(err)
	}
}
