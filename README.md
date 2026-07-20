# go-dns-resolver
Simple DNS resolver written in Go

## What and Why
This project provides a simple DNS resolver implementation in Go, allowing you to perform lookups by domain name, IP address, and reverse DNS lookup.

## Features

- **Lookup by domain name**: Resolve a domain name to its corresponding IP addresses.
- **Lookup by IP address**: Resolve an IP address to its corresponding domain name.
- **Reverse DNS lookup**: Perform a reverse DNS lookup on a given IP address.

## Installation
To use this resolver, run the following command:
```bash
go get github.com/miekg/dns
go build
```
This will download the required `github.com/miekg/dns` package and compile the resolver.

## Usage
To use the resolver, simply run the compiled executable and pass the desired lookup type and input as command-line arguments:
```bash
./go-dns-resolver lookup <domain_name> (or) <ip_address> (or) <ip_address> for reverse DNS lookup
```
For example:
```bash
./go-dns-resolver lookup example.com
./go-dns-resolver lookup 8.8.8.8
./go-dns-resolver reverse 8.8.8.8
```
## Build from Source
To build the resolver from source, run the following command:
```bash
make
```
This will compile the resolver and create an executable file named `go-dns-resolver`.

## Project Structure
The project structure is as follows:
```bash
go-dns-resolver/
Makefile
go.mod
dns.go
dns_test.go
.gitignore
```
## Dependencies
This project depends on the `github.com/miekg/dns` package for DNS resolution.

## License
This project is licensed under the MIT License.

## Contributing
Contributions are welcome! If you'd like to contribute, please submit a pull request or open an issue.

## Architecture
See [ARCHITECTURE.md](ARCHITECTURE.md) for a high-level overview of the project architecture.

## Tests
Unit tests for the resolver are located in `dns_test.go`. To run the tests, use the following command:
```bash
go test
```