package main

import (
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(">> ")
}

func DemoStdin() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed reading stadin: %s", err)
	}
	log.Printf("read %d from stdin: %s", len(input), input)
}

func DemoDevNull() {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatalf("failed opening nu.. device: %s", err)
	}
	defer devNull.Close()
	io.WriteString(devNull, "This is going nowhere\n")
}

func main() {
	io.WriteString(os.Stdout, "this is stdout\n")
	io.WriteString(os.Stderr, "This is stderr\n")
	DemoDevNull()
	DemoStdin()
}
