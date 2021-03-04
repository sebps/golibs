package connectors

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func copyConn(src net.Conn, targetHost string, targetPort int64) {
	dst, err := net.Dial("tcp", targetHost+":"+strconv.FormatInt(targetPort, 10))
	if err != nil {
		panic("Dial Error:" + err.Error())
	}

	done := make(chan struct{})

	go func() {
		defer src.Close()
		defer dst.Close()
		io.Copy(dst, src)
		done <- struct{}{}
	}()

	go func() {
		defer src.Close()
		defer dst.Close()
		io.Copy(src, dst)
		done <- struct{}{}
	}()

	<-done
	<-done
}

func TcpPiper(c *Connexion) {
	c.Protocol = "tcp"

	if len(c.SourceHost) == 0 {
		c.SourceHost = "localhost"
	}
	if len(c.TargetHost) == 0 {
		c.TargetHost = "localhost"
	}

	target := c.GetTarget()
	source := c.GetSource()

	log.Printf("TCP Redirection : \n%s://%s -> %s://%s\n", source.Scheme, source.Host, target.Scheme, target.Host)

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(w, req)
	})

	_, port, _ := net.SplitHostPort(source.Host)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("connection error:" + err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Error:", err)
			continue
		}
		copyConn(conn, c.TargetHost, c.TargetPort)
	}
}
