package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var url = flag.Bool("url", false, "Use URLEncoding instead of StdEncoding")

func data() []byte {
	data, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	return data
}

func encoding() *base64.Encoding {
	if *url {
		return base64.URLEncoding
	}
	return base64.StdEncoding
}

func main() {
	flag.Parse()
	var buffer bytes.Buffer
	enc := base64.NewEncoder(encoding(), io.MultiWriter(os.Stdout, &buffer))
	_, err := enc.Write(data())
	enc.Close()
	if err != nil {
		log.Fatalf("failed encoding: %s", err)
	}
	fmt.Println()
	dec := base64.NewDecoder(encoding(), &buffer)
	log.Println("decoding to stdout")
	io.Copy(os.Stdout, dec)
}
