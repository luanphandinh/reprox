package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var servers = map[string]string{
	"server1": "http://localhost:3000/",
	"server2": "http://localhost:3001/",
}

func main() {
	proxies := make(map[string]*httputil.ReverseProxy)

	for name, serverUrl := range servers {
		server, err := url.Parse(serverUrl)
		if err != nil {
			panic(err)
		}

		proxies[name] = httputil.NewSingleHostReverseProxy(server)
	}

	log.Println("Reverse Proxy server start in port 8080")

	http.HandleFunc("/", handler(proxies))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(proxies map[string]*httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyName := ""
		path := "/"
		paths := strings.SplitN(r.URL.Path, "/", 3)
		if len(paths) > 1 {
			proxyName = paths[1]
		}

		if len(paths) > 2 {
			path += paths[2]
		}

		if proxyName == "" {
			return
		}

		if proxy := proxies[proxyName]; proxy != nil {
			r.URL.Path = path
			proxy.ServeHTTP(w, r)
		}

		return
	}
}
