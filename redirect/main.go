package main

import (
	"github.com/sebps/golibs/redirect/connectors"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	var protocol string
	c := &connectors.Connexion{}

	for _, args := range os.Args[1:] {
		splitted := strings.Split(args, "=")
		if len(splitted) != 2 {
			log.Fatal("Wrong argument argument : " + args)
			os.Exit(-1)
		}
		name := splitted[0]
		value := splitted[1]

		switch name {
		case "sourceHost", "-sourceHost", "--sourceHost":
			c.SourceHost = value
		case "sourcePort", "-sourcePort", "--sourcePort":
			c.SourcePort, err = strconv.ParseInt(value, 10, 64)
		case "targetHost", "-targetHost", "--targetHost":
			c.TargetHost = value
		case "targetPort", "-targetPort", "--targetPort":
			c.TargetPort, err = strconv.ParseInt(value, 10, 64)
		case "protocol", "-protocol", "--protocol":
			if value == "tcp" || value == "http" {
				protocol = value
			}
		}
	}

	if c.SourcePort == 0 {
		log.Fatal("sourcePort argument required ")
		os.Exit(-1)
	}
	if c.TargetPort == 0 {
		log.Fatal("targetPort argument required ")
		os.Exit(-1)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	switch protocol {
	case "http":
		connectors.HttpConnector(c)
	case "tcp":
		connectors.TcpConnector(c)
	default:
		connectors.HttpConnector(c)
	}
}
