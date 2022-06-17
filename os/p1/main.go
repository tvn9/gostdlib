// Basic change file access permission by func Chmod
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// read cli-argument
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a file name, then press ENTER")
		return
	}

	fname := args[0]

	file, err := os.Open(fname) // For read access.
	if err != nil {
		log.Fatalf("fails to open file %s, error: %s\n", fname, err)
	}
	defer file.Close()

	// Set up a slice of byte with pre-define counts
	data := make([]byte, 500)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
