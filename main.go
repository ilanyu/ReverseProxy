package main

import (
	"log"
	"net/http"
	"bufio"
	"os"
	"strings"
)

var srv http.Server

func StartServer(bind string, remote string)  {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	h := &handle{reverseProxy: remote}
	srv.Addr = bind
	srv.Handler = h

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("ListenAndServe: ", err)
		}
	}()
}

func StopServer()  {
	if err := srv.Close() ; err != nil {
		log.Println(err)
	}
}

func main() {
	cmd := parseCmd()
	StartServer(cmd.bind, cmd.remote)
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		if strings.TrimSpace(str) == "stop" {
			log.Println("will stop server")
			StopServer()
			return
		}
	}
}
