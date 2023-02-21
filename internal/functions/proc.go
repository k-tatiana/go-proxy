package functions

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverCount = 0

// These constant is used to define all backend servers
const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
	PORT    = "1338"
)

// Balance returns one of the servers using round-robin algorithm
func GetProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[serverCount]
	serverCount++
	// reset the counter and start from the beginning
	if serverCount >= len(servers) {
		serverCount = 0
	}
	return server
}

// Serve a reverse proxy for a given url
func ServeReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)
	// Note that ServeHttp is non blocking & uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
