package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StringServer string

func (s StringServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("Received form  data: %v\n", r.Form)
	w.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello World"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	useRequest()
	simplePost()
}

func simplePost() {
	res, err := http.Post("http://localhost:7070",
		"applicaion/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}

func useRequest() {
	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "Radek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest("POST",
		"http://localhost:7070",
		strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded")

	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}
