// func Copy
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//Reading cli argument
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Enter a file1 file2 to copy")
		return
	}

	file1 := args[0]
	file2 := args[1]

	// Open existing file
	f1, err := os.Open(file1)
	if err != nil {
		fmt.Errorf("failed to open file %s\n", err)
	}
	defer f1.Close()

	// Create new file
	f2, err := os.Create(file2)
	if err != nil {
		fmt.Errorf("failed to create file %s\n", file2)
	}

	// Set up a strings reader
	if _, err := io.Copy(f2, f1); err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(file2)
	if err != nil {
		fmt.Errorf("failed to read %s\n", file2)
	}
	// Write file content to terminal
	os.Stdout.Write(data)
}
