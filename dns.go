// Package dns provides a simple DNS resolver implementation.
package dns

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

// Resolver represents a DNS resolver.
type Resolver struct {
	dns.Client
}

// LookupByName performs a DNS lookup by domain name.
func (r *Resolver) LookupByName(ctx context.Context, name string) (*dns.Msg, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(name), dns.TypeA)
	r.Client.Exchange(m, dns.Fqdn(name))
	return m, nil
}

// LookupByIP performs a DNS lookup by IP address.
func (r *Resolver) LookupByIP(ctx context.Context, ip net.IP) (*dns.Msg, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(ip.String()), dns.TypePTR)
	r.Client.Exchange(m, dns.Fqdn(ip.String()))
	return m, nil
}

// ReverseLookup performs a reverse DNS lookup.
func (r *Resolver) ReverseLookup(ctx context.Context, ip net.IP) (*dns.Msg, error) {
	m := new(dns.Msg)
	m.SetQuestion(net.IP(ip).String(), dns.TypePTR)
	r.Client.Exchange(m, net.IP(ip).String())
	return m, nil
}

// NewResolver returns a new DNS resolver instance.
func NewResolver() *Resolver {
	r := &Resolver{}
	r.Client.Net = "udp"
	return r
}

func main() {
	r := NewResolver()
	name := "example.com"
	ip := net.ParseIP("8.8.8.8")

	m, err := r.LookupByName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DNS record for %s: %s\n", name, m.Answer[0].(*dns.A).IP)

	m, err = r.LookupByIP(context.Background(), ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reverse DNS record for %s: %s\n", ip, m.Answer[0].(*dns.PTR).Ptr)

	m, err = r.ReverseLookup(context.Background(), ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reverse DNS record for %s: %s\n", ip, m.Answer[0].(*dns.PTR).Ptr)
}