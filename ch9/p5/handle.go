// Handling HTTP requests
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User POST")
		}
	})

	// separate handler
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the contact page")
	})

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the Clothes page")
	})

	// Default server
	http.ListenAndServe(":8080", mux)
}
