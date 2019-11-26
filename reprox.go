package reprox

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Reprox struct {
	proxies map[string]*httputil.ReverseProxy
}

func NewReprox() *Reprox {
	return &Reprox{
		proxies: map[string]*httputil.ReverseProxy{},
	}
}

func (r *Reprox) Register(name string, path string) {
	remote, err := url.Parse(path)
	if err != nil {
		panic(err)
	}

	r.proxies[name] = httputil.NewSingleHostReverseProxy(remote)
}

func (rp *Reprox) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyName := ""
		path := "/"
		paths := strings.SplitN(r.URL.Path, "/", 3)
		if len(paths) > 1 {
			proxyName = paths[1]
		}

		if proxyName == "" {
			return
		}

		if len(paths) > 2 {
			path += paths[2]
		}

		if proxy := rp.proxies[proxyName]; proxy != nil {
			r.URL.Path = path
			proxy.ServeHTTP(w, r)
		}

		return
	}
}
