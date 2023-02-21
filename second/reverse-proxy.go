package second

import (
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	FIRST_HOST  = "127.0.0.1:9090"
	SECOND_HOST = "127.0.0.1:8081"
)

func SecondProxy() *httputil.ReverseProxy {

	new := newUrl(FIRST_HOST)
	proxy := httputil.NewSingleHostReverseProxy(new)
	return proxy
}

func newUrl(address string) *url.URL {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	newUrl, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	return newUrl
}
