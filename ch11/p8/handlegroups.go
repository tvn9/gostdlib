// Breaking HTTP handlers into groups
// this example demonstrates how the handlers could be separated into modules.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func restModule() http.Handler {
	// Separate Mux for all REST
	restApi := http.NewServeMux()
	restApi.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"name":"john"}]`)
	})
	return restApi
}

func uiModule() http.Handler {
	// Separate Mux for all UI
	ui := http.NewServeMux()
	ui.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<html><body><h1>Hello from UI!</h1></body></html>`)
	})
	return ui
}

func main() {
	log.Println("Starting server...")
	// Adding to mani Mux
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", restModule()))
	mux.Handle("/ui/", http.StripPrefix("/ui", uiModule()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
