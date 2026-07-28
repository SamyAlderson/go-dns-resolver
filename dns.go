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

// NewResolver returns a new DNS resolver instance.
func NewResolver() *Resolver {
	r := &Resolver{}
	r.Client.Net = "udp"
	return r
}

// LookupByName performs a DNS lookup by domain name.
func (r *Resolver) LookupByName(ctx context.Context, name string) (*dns.Msg, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(name), dns.TypeA)
	return r.Client.Exchange(m, dns.Fqdn(name))
}

// LookupByIP performs a DNS lookup by IP address.
func (r *Resolver) LookupByIP(ctx context.Context, ip net.IP) (*dns.Msg, error) {
	if !ip.IsGlobalUnicast() {
		return nil, fmt.Errorf("invalid IP address: %s", ip)
	}
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(ip.String()), dns.TypePTR)
	return r.Client.Exchange(m, dns.Fqdn(ip.String()))
}

// ReverseLookup performs a reverse DNS lookup.
func (r *Resolver) ReverseLookup(ctx context.Context, ip net.IP) (*dns.Msg, error) {
	if !ip.IsGlobalUnicast() {
		return nil, fmt.Errorf("invalid IP address: %s", ip)
	}
	m := new(dns.Msg)
	m.SetQuestion(net.IP(ip).String(), dns.TypePTR)
	return r.Client.Exchange(m, net.IP(ip).String())
}

func main() {
	r := NewResolver()
	name := "example.com"
	ip := net.ParseIP("8.8.8.8")

	if m, err := r.LookupByName(context.Background(), name); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("DNS record for %s: %s\n", name, m.Answer[0].(*dns.A).IP)
	}

	if m, err := r.LookupByIP(context.Background(), ip); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Reverse DNS record for %s: %s\n", ip, m.Answer[0].(*dns.PTR).Ptr)
	}

	if m, err := r.ReverseLookup(context.Background(), ip); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Reverse DNS record for %s: %s\n", ip, m.Answer[0].(*dns.PTR).Ptr)
	}
}