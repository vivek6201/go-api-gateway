package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/vivek6201/go-api-gateway/internals/utils"
)

// ForwardRequest proxies to targetBase. stripPrefix is the gateway path prefix to remove
// from the outgoing URL path (e.g. "/users" so /users/list becomes /list upstream).
func ForwardRequest(targetBase string, stripPrefix string, w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(targetBase)

	if err != nil {
		utils.JSONError(w, http.StatusBadGateway, "invalid target URL")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	orig := proxy.Director
	proxy.Director = func(req *http.Request) {
		if stripPrefix != "" {
			req.URL.Path = stripGatewayPrefix(req.URL.Path, stripPrefix)
			req.URL.RawPath = ""
		}
		orig(req)
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, _ error) {
		utils.JSONError(w, http.StatusInternalServerError, "upstream request failed")
	}

	proxy.Transport = &http.Transport{
		ResponseHeaderTimeout: 3 * time.Second,
	}

	proxy.ServeHTTP(w, r)
}

func stripGatewayPrefix(path, prefix string) string {
	if path == prefix || path == prefix+"/" {
		return "/"
	}
	if strings.HasPrefix(path, prefix+"/") {
		suffix := strings.TrimPrefix(path[len(prefix):], "/")
		if suffix == "" {
			return "/"
		}
		return "/" + suffix
	}
	return path
}
