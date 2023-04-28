package service

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

type Saml struct {
	saml *samlsp.Middleware
}

type SamlIFace interface {
	GetSamlMiddleware() *samlsp.Middleware
	GetUsernameFromContext(ctx context.Context) string
	GetEmailFromContext(ctx context.Context) string
	GetIdFromContext(ctx context.Context) string
}

func initializeSaml(log LoggerIFace) SamlIFace {
	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	// the idpMetadataURL can be hardcoded since we are only supporting Elon's IdP instance
	// this url is for testing purposes only—we will be using Elon's metadata in prod.
	// todo(@jumar): replace this with Elon's Metadata URL
	idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
	if err != nil {
		log.Fatal(err.Error())
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	rootURL, err := url.Parse("http://localhost:3000")
	if err != nil {
		log.Fatal(err.Error())
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	return &Saml{
		saml: samlSP,
	}
}

func (s *Saml) GetSamlMiddleware() *samlsp.Middleware {
	return s.saml
}

func (s *Saml) GetUsernameFromContext(ctx context.Context) string {
	return samlsp.AttributeFromContext(ctx, "username")
}

func (s *Saml) GetIdFromContext(ctx context.Context) string {
	return samlsp.AttributeFromContext(ctx, "id")
}

func (s *Saml) GetEmailFromContext(ctx context.Context) string {
	return samlsp.AttributeFromContext(ctx, "email")
}
