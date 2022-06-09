// Creating HTTP middleware layer
// Modern applications with web UI or REST API usually use the middleware mechanism to log the activity,
// or guard the security of the given interface. In the example, the implemetation of such a middleware
// layer will be presented.
package main

import (
	"io"
	"net/http"
)

func main() {
	// Secured API
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `[{"id":"1","login":"ffghi"}]`)
	}))
	http.ListenAndServe(":8080", mux)
}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-Auth")
		if sec != "authenticated" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(w, r) // use the handler
	}
}
