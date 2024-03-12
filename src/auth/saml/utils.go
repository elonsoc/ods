package saml

import (
	"net"
	"net/url"
	"strings"
)

// getDomainFromURI formats a domain string into a proper
// domain name to be inlayed into a cookie.
func getDomainFromURI(domain string) (string, error) {
	if strings.ToLower(domain[:4]) == "http" {
		u, err := url.Parse(domain)
		if err != nil {
			return "", err
		}
		return u.Hostname(), nil
	}
	// the provided domain is not a URL, so it should be a hostname
	domain, _, err := net.SplitHostPort(domain)
	if err != nil {
		return "", err

	}

	return domain, nil

}
