package connectors

import (
	"log"
	"net/url"
	"strconv"
)

type Connexion struct {
	SourceHost string
	SourcePort int64
	TargetHost string
	TargetPort int64
	Protocol   string
}

func (c Connexion) GetSource() *url.URL {
	source, err := url.Parse(c.Protocol + "://" + c.SourceHost + ":" + strconv.FormatInt(c.SourcePort, 10))
	if err != nil {
		log.Fatal(err)
	}

	return source
}

func (c Connexion) GetTarget() *url.URL {
	target, err := url.Parse(c.Protocol + "://" + c.TargetHost + ":" + strconv.FormatInt(c.TargetPort, 10))
	if err != nil {
		log.Fatal(err)
	}

	return target
}
