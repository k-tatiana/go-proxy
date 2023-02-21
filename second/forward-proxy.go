package second

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

var hopHeaders = []string{
	"Connection",
	"Proxy-Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailer",
	"Transfer-Encoding",
	"Upgrade",
}

type forwardProxy struct {
}

func removeConnectionHeaders(h http.Header) {
	for _, f := range h["Connection"] {
		for _, sf := range strings.Split(f, ",") {
			if sf = strings.TrimSpace(sf); sf != "" {
				h.Del(sf)
			}
		}
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

func (fp *forwardProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// The "Host:" header is promoted to Request.Host and is removed from
	// request.Header by net/http, so we print it out explicitly.
	log.Println(request.RemoteAddr, "\t\t", request.Method, "\t\t", request.URL, "\t\t Host:", request.Host)
	log.Println("\t\t\t\t\t", request.Header)
	client := &http.Client{}
	request.RequestURI = ""
	for _, h := range hopHeaders {
		request.Header.Del(h)
	}
	removeConnectionHeaders(request.Header)
	clientIP, _, _ := net.SplitHostPort(request.RemoteAddr)
	appendHostToXForwardHeader(request.Header, clientIP)
	resp, err := client.Do(request)
	if err != nil {
		http.Error(writer, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
	}
	defer resp.Body.Close()

	log.Println(request.RemoteAddr, " ", resp.Status)

	for _, h := range hopHeaders {
		request.Header.Del(h)
	}
	removeConnectionHeaders(resp.Header)

	for k, vv := range resp.Header {
		for _, v := range vv {
			writer.Header().Add(k, v)
		}
	}
	writer.WriteHeader(resp.StatusCode)
	io.Copy(writer, resp.Body)
}

func ReverseProxy() *forwardProxy {
	address := "127.0.0.1:9999"
	proxy := &forwardProxy{}
	log.Println("Starting proxy server on", address)
	return proxy
}
