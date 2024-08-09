// file: pkg/snettool/dnslookup.go
package snettool

import (
    "net"
)

// DNSLookup performs a DNS lookup for the specified hostname and returns its IP addresses.
func DNSLookup(hostname string) ([]string, error) {
    ips, err := net.LookupIP(hostname)
    if err != nil {
        return nil, err
    }

    var ipStrs []string
    for _, ip := range ips {
        ipStrs = append(ipStrs, ip.String())
    }

    return ipStrs, nil
}
