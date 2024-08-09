// file: pkg/snettool/whoislookup.go
package snettool

import (
	"github.com/likexian/whois"
)

// WhoisLookup performs a WHOIS lookup for the specified domain and returns the result.
func WhoisLookup(domain string) (string, error) {
	result, err := whois.Whois(domain)
	if err != nil {
		return "", err
	}
	return result, nil
}
