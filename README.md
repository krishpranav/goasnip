# goasnip
goasnip retrieves all IPs of a target organization—used for attack surface mapping in reconnaissance.

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation
```
git clone https://github.com/krishpranav/goasnip
cd goasnip
go build goasnip.go
./goasnip
```

# Options
```
Usage:
  -t string
        Domain or IP address (Required)
  -p string
        Print results to console
```

# Example
```
$ goasnip -t google.com -p

[?] ASN: "15169" ORG: "GOOGLE, US"
8.8.4.0/24
--- snip ---
[:] Writing 616 CIDRs to file...
[:] Converting to IPs...
8.8.8.1
--- snip ---
[:] Writing 14725936 IPs to file...
[!] Done.
```
