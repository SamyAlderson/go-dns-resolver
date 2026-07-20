# go-dns-resolver

I built this because I got tired of using the built-in `net.LookupIP()` and `net.LookupAddr()` functions in Go, only to realize they return incomplete results or time out on me. 

## What I wanted

A simple DNS resolver that can look up domain names, IP addresses, and do reverse DNS lookups without the hassle of dealing with the official Go libraries.

## Features

- `LookupDomain`: does a forward DNS lookup for a given domain name
- `LookupIP`: does a reverse DNS lookup for a given IP address
- `LookupAddr`: does a forward DNS lookup for a given IP address

I know what you're thinking: "Why not just use `net.LookupIP()` and `net.LookupAddr()`?" Well, those functions are great, but they don't always return the full set of results. Sometimes they'll return only the A records, and you might miss the MX or NS records. Sometimes they'll time out, and you'll be stuck waiting for the timeout to expire. Not cool.

## Why not just use the `github.com/miekg/dns` package?

Because it's a bit much for what I needed. Don't get me wrong, Miek Gieben's `dns` package is awesome, but it's got a lot of features I don't need. I just wanted a simple DNS resolver that can do forward and reverse lookups. This project does that.

## The Problem with DNS

People often complain about DNS being broken or slow, but the real problem is that the APIs for DNS are just as messy as the underlying protocol. You gotta deal with timeouts and incomplete results and all sorts of other nonsense. This project tries to abstract that away a bit.

## Contributing

If you've got a better way to do DNS lookups in Go, open a PR. I'd love to see it.