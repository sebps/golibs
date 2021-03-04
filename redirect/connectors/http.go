package connectors

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func HttpConnector(c *Connexion) {
	c.Protocol = "http"

	if len(c.SourceHost) == 0 {
		c.SourceHost = "localhost"
	}
	if len(c.TargetHost) == 0 {
		c.TargetHost = "localhost"
	}

	target := c.GetTarget()
	source := c.GetSource()

	log.Printf("HTTP Redirection : \n%s://%s -> %s://%s\n", source.Scheme, source.Host, target.Scheme, target.Host)

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(w, req)
	})

	_, port, _ := net.SplitHostPort(source.Host)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
