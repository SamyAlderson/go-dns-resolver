package main

import (
	"testing"
	"github.com/miekg/dns"
)

func TestLookupDomain(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("example.com.", dns.TypeA)
	res, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "8.8.8.8:53")
	if err != nil {
		t.Errorf("dns lookup failed: %v", err)
	}
	if len(res.Answer) == 0 {
		t.Errorf("no answer for domain lookup")
	}
}

func TestLookupIP(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("8.8.8.8", dns.TypePTR)
	res, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "8.8.8.8:53")
	if err != nil {
		t.Errorf("dns lookup failed: %v", err)
	}
	if len(res.Answer) == 0 {
		t.Errorf("no answer for IP lookup")
	}
}

func TestReverseLookup(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("8.8.8.8.in-addr.arpa.", dns.TypePTR)
	res, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "8.8.8.8:53")
	if err != nil {
		t.Errorf("dns lookup failed: %v", err)
	}
	if len(res.Answer) == 0 {
		t.Errorf("no answer for reverse lookup")
	}
}

func TestInvalidInput(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("", dns.TypeA)
	_, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "8.8.8.8:53")
	if err == nil {
		t.Errorf("expecting error for invalid input")
	}
}