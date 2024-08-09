// file: pkg/snettool/dnslookup_test.go
package snettool

import (
    "testing"
)

func TestDNSLookup(t *testing.T) {
    result, err := DNSLookup("google.com")
    if err != nil {
        t.Errorf("DNSLookup failed: %v", err)
    }
    if len(result) == 0 {
        t.Errorf("Expected non-empty result, got %v", result)
    }
}
