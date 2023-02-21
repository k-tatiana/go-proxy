package main

import (
	"log"
	"net/http"

	"github.com/tin-proxy/second"
)

func main() {
	// start server
	proxy := second.ReverseProxy()
	if err := http.ListenAndServe("127.0.0.1:9999", proxy); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
