package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"flag"
)

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
}

func main() {
	bind := flag.String("l", "0.0.0.0:8888", "listen on ip:port")
	remote := flag.String("r", "http://idea.lanyus.com:80", "reverse proxy addr")
	flag.Parse()
	log.Printf("Listening on %s, forwarding to %s", *bind, *remote)
	h := &handle{reverseProxy: *remote}
	err := http.ListenAndServe(*bind, h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
