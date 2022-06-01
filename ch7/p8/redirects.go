// Handling HTTP redirects
package main

import (
	"fmt"
	"net/http"
)

const addr = "localhost:7070"

type RedirecServer struct {
	redirectCount int
}

func (s *RedirecServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.redirectCount++
	fmt.Println("Received header: " + r.Header.Get("Known-redirects"))
	http.Redirect(w, r, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func main() {
	s := http.Server{
		Addr:    addr,
		Handler: &RedirecServer{redirectCount: 0},
	}
	go s.ListenAndServe()

	client := http.Client{}
	redirectCount := 0

	// If the count of reriects is reached
	// than return error.
	client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		if redirectCount > 2 {
			return fmt.Errorf("too many redirects")
		}
		r.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}
}
