package main

import (
	// Generate a privae key
	"crypto/rsa"
	// Ceate a TLS connection which encrypts data to ensure security
	"crypto/tls"
	// Creates X.509 certificates, which are then used in conjunction 
	// with the sessioncert file to sign sessions.
	"crypto/x509"
	// 
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

// Variable containing our Idp's xml metadata URL to be queried for authenticating users
var metdataurl="https://idp.elon.edu/idp/shibboleth" 
// Variable containing public key (.pem file) from RSA 
var sessioncert="./sessioncert"
// Variable containing private key (.pem file) from RSA 
var sessionkey="./sessionkey"
var serverkey="./serverkey"  //Server TLS 
var servercert="./servercert"
 // Variable containing url of this service
var serverurl="https://localhost" 
// Variable contraining server url to uniquely identify your service for  the IDP
var entityId=serverurl
// ????
var listenAddr="0.0.0.0:443"

// Handler for HTTP request, takes request writer and response parameters
// Returns http.ResponseWriter with text
func hello(w http.ResponseWriter, r *http.Request) {
	// Checks to see if session is already avaliable on r's context				????
	s := samlsp.SessionFromContext(r.Context())
	// If not return without making any action
	if s == nil {
		return
	}
	// Otherwise, creates samlsp session from r's context using URL containing XML metadata
	sa, ok := s.(samlsp.SessionWithAttributes)
	if !ok {
		return
	}
	fmt.Fprintf(w, "Token contents, %+v!", sa.GetAttributes())
}

func main() {
	// Creates key pair and sores into keyPair using TLS
	keyPair, err := tls.LoadX509KeyPair(sessioncert,sessionkey)
	// Checks for error
	panicIfError(err)
	// Parses cert from first key with x509 and stores in keyPair
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	panicIfError(err)
	// Uses url.Parse to convert the metadata in our URL
	idpMetadataURL, err := url.Parse(metdataurl)
	panicIfError(err)
	// ' ' ' to convert serverurl server url 
	rootURL, err := url.Parse(serverurl)
	panicIfError(err)
	// Variables are used to specify our samlSP 
	samlSP, _ := samlsp.New(samlsp.Options{
		// rootURL is the root of our domain used as identifier for authentification
		URL:            *rootURL,
		// keyPair set to an rsa.PrivateKey which allows for it to be used 
		// with the certificate leaf entityId from IDPMetadataURL 
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadataURL: https://idp.elon.edu/idp/shibboleth,
		EntityID:   	  entityId,
	})
	// Handle requests from the /hello endpoint
	app := http.HandlerFunc(hello)
	// 
	http.Handle("/hello", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	panicIfError(http.ListenAndServeTLS(listenAddr,servercert,serverkey, nil))
}

func panicIfError(err error){
	if err!=nil{
		panic(err)
	}
}


