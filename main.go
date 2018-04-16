package main

import (
	"net/http"
	"log"
)

var cmd Cmd
var srv http.Server

func StartServer(bind string, remote string)  {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	h := &handle{reverseProxy: remote}
	srv.Addr = bind
	srv.Handler = h
	//go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("ListenAndServe: ", err)
		}
	//}()
}

func StopServer()  {
	if err := srv.Shutdown(nil) ; err != nil {
		log.Println(err)
	}
}

func main() {
	cmd = parseCmd()
	StartServer(cmd.bind, cmd.remote)
}
