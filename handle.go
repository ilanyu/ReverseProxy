package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"log"
	"net"
	"context"
	"time"
	"io/ioutil"
)

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		if addr == "idea.lanyus.com:80" {
			resp, err := http.Get("http://119.29.29.29/d?dn=idea.lanyus.com")
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
			res, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			addr = string(res) + ":80"
		}

		return dialer.DialContext(ctx, network, addr)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
}
