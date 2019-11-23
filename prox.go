package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	serverUrls := []string{
		"http://localhost:3000/",
		"http://localhost:3001/",
	}

	proxies := make([]*httputil.ReverseProxy, 2)

	for i, serverUrl := range serverUrls {
		server, err := url.Parse(serverUrl)
		if err != nil {
			panic(err)
		}

		proxies[i] = httputil.NewSingleHostReverseProxy(server)
	}

	log.Println("Reverse Proxy server start in port 8080")

	http.HandleFunc("/", handler(proxies))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(proxies []*httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		proxy := proxies[0]
		if r.URL.Path == "/2" {
			proxy = proxies[1]
		}

		proxy.ServeHTTP(w, r)
	}
}
