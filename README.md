# OSCP-go

This repository contains a HTTP client (in Go) for OSCP (Open Smart Charging Protocol), which was generated using
the [oapi-codegen](https://github.com/deepmap/oapi-codegen#overview) generator.

It also contains an OpenAPI specification for OSCP, which is used to generate the client and can be used to generate
a compatible HTTP API server.

## Installation

```bash
 go get github.com/ChargePi/oscp-go@latest"
```

## Usage

```go
import "github.com/ChargePi/oscp/<oscp-version>"
```