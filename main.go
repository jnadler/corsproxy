package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var ListenPort = "9200"
var RemoteUrl = "http://elasticsearch:9200/"

// Add CORS Headers
func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		handler.ServeHTTP(w, r)
	})
}

// Replace the hostname
func setDestinationHost(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = r.URL.Host
		handler.ServeHTTP(w, r)
	})
}

// Run a SingleHostReverseProxy to RemoteUrl on local port ListenPort
func Proxy(remoteUrl string) http.Handler {
	serverUrl, err := url.Parse(remoteUrl)
	if err != nil {
		log.Fatal("URL failed to parse")
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(serverUrl)
	singleHosted := setDestinationHost(reverseProxy)
	return addCORSHeaders(singleHosted)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		println("usage: corsproxy <listen-port> <remote-url>")
		os.Exit(1)
	}

	ListenPort = args[1]
	RemoteUrl = args[2]

	println("Starting proxy for " + RemoteUrl + " on local port " + ListenPort)
	proxy := Proxy(RemoteUrl)
	http.ListenAndServe(":"+ListenPort, proxy)
}
