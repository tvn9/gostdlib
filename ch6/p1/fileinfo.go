// Getting file information

package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s\n", fi.Name())
	fmt.Printf("Is directory: %t\n", fi.IsDir())
	fmt.Printf("File size: %d\n", fi.Size())
	fmt.Printf("File mode: %v\n", fi.Mode())
	fmt.Printf("File mode: %v\n", fi.ModTime())
}
