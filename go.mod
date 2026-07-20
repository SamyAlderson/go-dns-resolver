go.mod
```go
module go-dns-resolver

go 1.19

require (
	github.com/miekg/dns v1.4.3
)

replace (
	github.com/miekg/dns v1.4.3 => github.com/miekg/dns v1.4.3
)

require (
	github.com/miekg/dns/v2 v2.1.0
)
```