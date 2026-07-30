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
	if res.Answer[0].Header().Rcode != dns.RcodeSuccess {
		t.Errorf("unexpected rcode: %v", res.Answer[0].Header().Rcode)
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
	if res.Answer[0].Header().Rcode != dns.RcodeSuccess {
		t.Errorf("unexpected rcode: %v", res.Answer[0].Header().Rcode)
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
	if res.Answer[0].Header().Rcode != dns.RcodeSuccess {
		t.Errorf("unexpected rcode: %v", res.Answer[0].Header().Rcode)
	}
}

func TestInvalidInput(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("", dns.TypeA)
	_, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "8.8.8.8:53")
	if err == nil {
		t.Errorf("expecting error for invalid input")
	}
	if err.Error() != "question section empty" {
		t.Errorf("unexpected error message: %v", err.Error())
	}
}

func TestExchangeError(t *testing.T) {
	dns.DefaultMsg = &dns.Msg{}
	dns.DefaultMsg.SetQuestion("example.com.", dns.TypeA)
	_, err := dns.DefaultMsg.Exchange(dns.DefaultMsg, "invalid:53")
	if err == nil {
		t.Errorf("expecting error for invalid server address")
	}
	if err.Error() != "dial invalid:53: lookup invalid:53: no address associated with hostname" {
		t.Errorf("unexpected error message: %v", err.Error())
	}
}