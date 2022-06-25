package main

import (
	"bytes"
	"encoding/base32"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var hex = flag.Bool("hex", false, "Use HexEncoding instead of StdEncoding")

func data() []byte {
	data, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	return data
}

func encoding() *base32.Encoding {
	if *hex {
		return base32.HexEncoding
	}
	return base32.StdEncoding
}

func main() {
	flag.Parse()
	var buffer bytes.Buffer
	enc := base32.NewEncoder(encoding(), io.MultiWriter(os.Stdout, &buffer))
	log.Println("encoding to stdout")
	_, err := enc.Write(data())
	enc.Close()
	if err != nil {
		log.Fatalf("failed encoding: %s", err)
	}
	fmt.Println()
	dec := base32.NewDecoder(encoding(), &buffer)
	log.Println("decoding to stdout")
	io.Copy(os.Stdout, dec)
}
