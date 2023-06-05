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
}

func initializeSaml(log LoggerIFace, idpURL, certPath, keyPath string) SamlIFace {
	keyPair, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadataURL, err := url.Parse(idpURL)
	if err != nil {
		panic(err) // TODO handle error
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	// idpMetadata.Organization = &saml.Organization{
	// 	OrganizationNames: []saml.LocalizedName{
	// 		{Value: "Elon Society of Computing CS Project Team", Lang: "en"},
	// 	},
	// 	OrganizationURLs: []saml.LocalizedURI{
	// 		{Value: "https://ods.elon.edu", Lang: "en"},
	// 	},
	// 	OrganizationDisplayNames: []saml.LocalizedName{
	// 		{Value: "ESC CSPT", Lang: "en"},
	// 	},
	// }

	rootURL, err := url.Parse("https://ods.elon.edu")
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		EntityID:    rootURL.String(),
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
