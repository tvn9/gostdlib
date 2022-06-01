// Connecting to the remote server
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type StringServer string

func (s StringServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("HELLO GOPHER!\n"),
	}
}

const addr = "localhost:7070"

func main() {
	// 1. Create a server with on localhost port :7070
	s := createServer(addr)
	go s.ListenAndServe()

	// Connec with plain TCP with the localhost:7070
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(conn)

	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
