// Resolving form variables
package main

import (
	"fmt"
	"net/http"
)

type StringServer string

func (s StringServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", r.Form)
	r.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", r.Form)
	fmt.Println("Param1 is : " + r.Form.Get("param1"))
	w.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello World"),
	}
}

func main() {
	s := createServer(":8080")
	fmt.Println("Server is starting...")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
