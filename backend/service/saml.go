package service

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/elonsoc/saml"
	"github.com/elonsoc/saml/samlsp"
)

type Saml struct {
	saml *samlsp.Middleware
}

type SamlIFace interface {
	GetSamlMiddleware() *samlsp.Middleware
}

func initializeSaml(log LoggerIFace, idpURL, webURL, spURL, certPath, keyPath string) SamlIFace {
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

	rootURL, err := url.Parse(spURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	options := samlsp.Options{
		URL:               *rootURL,
		EntityID:          rootURL.String(),
		Key:               keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:       keyPair.Leaf,
		AllowIDPInitiated: false,
		IDPMetadata:       idpMetadata,
		Organization: &saml.Organization{
			OrganizationNames: []saml.LocalizedName{
				{Value: "Elon Society of Computing CS Project Team", Lang: "en"},
			},
			OrganizationURLs: []saml.LocalizedURI{
				{Value: "https://ods.elon.edu", Lang: "en"},
			},
			OrganizationDisplayNames: []saml.LocalizedName{
				{Value: "ESC CSPT", Lang: "en"},
			},
		},
		ContactPerson: &saml.ContactPerson{
			ContactType:    "technical",
			Company:        "Elon University",
			GivenName:      "Elon Society of Computing CS Project Team",
			EmailAddresses: []string{"hello@elonsoc.org"},
		},
		DefaultRedirectURI: "ods.elon.edu",
	}

	samlSP, _ := samlsp.New(options)

	samlSP.Session = samlsp.CookieSessionProvider{
		Name:     "escSamlCookie",
		Domain:   webURL,
		HTTPOnly: true,
		Secure:   options.URL.Scheme == "https",
		SameSite: options.CookieSameSite,
		Codec:    samlsp.DefaultSessionCodec(options),
	}

	return &Saml{
		saml: samlSP,
	}
}

func (s *Saml) GetSamlMiddleware() *samlsp.Middleware {
	return s.saml
}
