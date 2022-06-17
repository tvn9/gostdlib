package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader, with some random data 1\n")
	r2 := strings.NewReader("second reader, with some random data 2\n")
	buf := make([]byte, 10)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ...reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}
