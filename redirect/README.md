# Basic lib to implicitely redirect an incoming connexion to a target endpoint

## Protocol supported
- http 
- tcp

## Lib usage

```go
package main

import (
	"github.com/sebpsdev/golibs/redirect/connectors"
)

func main() {
	c := &connectors.Connexion{
		SourceHost: "127.0.0.1",
		SourcePort: 3000,
		TargetHost: "com.example.endpoint",
		TargetPort: 5000,
	}

	connectors.HttpPiper(c)
}
```

## Server Usage
go run main.go --sourceHost=127.0.0.1 --sourcePort=3000 --targetHost=com.example.endpoint --targetPort=5000 --protocol=http

## Required arguments
- sourcePort
- targetPort 

## Optional arguments
- sourceHost
- targetHost
- protocol

## Default values
- sourceHost: localhost
- targetHost: localhost
- protocol: http