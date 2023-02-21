package main

import (
	"log"
	"net/http"

	"github.com/tin-proxy/internal/functions"
	"github.com/tin-proxy/internal/handlers"
)

func main() {
	// start server
	http.HandleFunc("/", handlers.LoadBalancer)
	log.Fatal(http.ListenAndServe(":"+functions.PORT, nil))
}
