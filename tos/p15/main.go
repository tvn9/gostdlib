// ReadFile reads the named file and returns the contents. A successful call return err == nil,
// not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// get cli input
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a file name")
		return
	}

	fName := args[0]

	// Reading file content using tos.ReadFile
	data, err := os.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}

	// Writing file content to tos.Stdout
	os.Stdout.Write(data)
}
