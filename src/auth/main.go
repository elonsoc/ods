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
)

func initialize(certPath, keyPath, idpURL, spURL string) chi.Router {
	// startInitialization := time.Now()
	svc := common.NewService("", "", "", "", "", "")
	s := saml.InitializeSaml(svc, idpURL, spURL, spURL, certPath, keyPath)

	t := token.NewTokenServicer(svc.IMDb)

	smw := s.Middleware()

	r := chi.NewRouter()
	r.Mount("/saml", saml.SetSamlEndpoint(spURL, svc, t, smw))

	return r
}
func main() {
	servicePort := flag.String("service_port", os.Getenv("SERVICE_PORT"), "the port to run the service on")
	samlCertPath := flag.String("saml_cert_path", os.Getenv("SAML_CERT_PATH"), "location of service cert")
	samlKeyPath := flag.String("saml_key_path", os.Getenv("SAML_KEY_PATH"), "location of service key")
	idpURL := flag.String("idp_url", os.Getenv("IDP_URL"), "url of identity provider")
	spURL := flag.String("sp_url", os.Getenv("SP_URL"), "url of the hosted service provider")
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

	err := http.ListenAndServe(fmt.Sprintf(":%s", *servicePort),
		initialize(*samlCertPath, *samlKeyPath, *idpURL, *spURL))
	if err != nil {
		fmt.Println(err)
	}
}
