// Creating the HTTP Server
package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct {
}

func (s SimpleHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	fmt.Println("String HTTP server on port 8080")
	// Eventually you can use http.ListenAdnServer(":8080", SimpleHTTP{})
	s := &http.Server{
		Addr:    ":8080",
		Handler: SimpleHTTP{},
	}
	s.ListenAndServe()
}
