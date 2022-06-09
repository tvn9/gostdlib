// Serving secured HTTP content
// The example describes the simplest way of creating the HTTP server, which serves the content via TLS/SSL layer.

package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	fmt.Println("Starting HTTP server on port 8080")
	// Eventually you can use
	// http.LisenAndServe(":8080", SimpleHTTP{})
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	if err := s.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		panic(err)
	}
}
