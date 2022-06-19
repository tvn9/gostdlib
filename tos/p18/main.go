package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// get cli-input arguments
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Enter a name to create a new file")
		return
	}
	fName := args[0]

	// Create a new file
	file, err := os.Create(fName)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("File already exist - Error: %s\n", err)
	}

	// Writing some data to the created file
	if _, err := file.Write([]byte("Hello Gopher!\n")); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
