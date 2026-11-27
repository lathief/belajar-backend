package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Products meneruskan /api/products ke /products pada backend lokal.
func Products() http.Handler {
	target := &url.URL{Scheme: "http", Host: "127.0.0.1:9000"}
	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(target)
		},
	}
	return http.StripPrefix("/api", proxy)
}
