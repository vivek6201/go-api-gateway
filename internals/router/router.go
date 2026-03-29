package router

import (
	"net/http"

	"github.com/vivek6201/go-api-gateway/internals/config"
	"github.com/vivek6201/go-api-gateway/internals/proxy"
)

func NewRouter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target, prefix, ok := config.RouteForPath(r.URL.Path)
		if !ok {
			http.NotFound(w, r)
			return
		}
		proxy.ForwardRequest(target, prefix, w, r)
	})
}
