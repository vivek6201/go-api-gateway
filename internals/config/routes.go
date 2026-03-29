package config

import "strings"

// Routes maps a path prefix to a backend base URL. Any request whose path equals
// the prefix or starts with prefix+"/" is forwarded to that backend (longest prefix wins).
var Routes = map[string]string{
	"/users":  "http://localhost:8081",
	"/orders": "http://localhost:8082",
}

// RouteForPath returns the backend base URL and the matched gateway prefix (used to
// strip that prefix before forwarding so e.g. /users/list hits the upstream as /list).
func RouteForPath(path string) (target string, matchedPrefix string, ok bool) {
	var bestPrefix string
	var bestTarget string
	for prefix, tgt := range Routes {
		if !matchesPrefix(path, prefix) {
			continue
		}
		if len(prefix) > len(bestPrefix) {
			bestPrefix = prefix
			bestTarget = tgt
		}
	}
	if bestPrefix == "" {
		return "", "", false
	}
	return bestTarget, bestPrefix, true
}

func matchesPrefix(path, prefix string) bool {
	if path == prefix {
		return true
	}
	return strings.HasPrefix(path, prefix+"/")
}
