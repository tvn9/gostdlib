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
	buf := make([]byte, 100)

	// buf is used here...
	nb, err := io.CopyBuffer(f2, f1, buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of byte written %v\n", nb)

	data, err := os.ReadFile(file2)
	if err != nil {
		log.Fatal(err)
	}

	// Writing file content to terminal
	os.Stdout.Write(data)
}
