package handlers

import (
	"log"
	"net/http"

	fn "github.com/tin-proxy/internal/functions"
)

// Given a request send it to the appropriate url
func LoadBalancer(res http.ResponseWriter, req *http.Request) {
	// Get address of one backend server on which we forward request
	url := fn.GetProxyURL()
	// Log the request
	log.Printf("proxy_url: %s\n", url)
	// Forward request to original request
	fn.ServeReverseProxy(url, res, req)
}
