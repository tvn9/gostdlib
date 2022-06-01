// Creating a HTTP request
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type StringServer string

func (s StringServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("Received form data: %v\n", r.Form)
	fmt.Printf("Received header: %v\n", r.Header)
	w.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	form := url.Values{}
	form.Set("id", "5")
	form.Set("name", "Wolfgang")
	req, err := http.NewRequest(http.MethodPost, "http://localhost:7070",
		strings.NewReader(form.Encode()))

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "applicaion/x-www-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}
