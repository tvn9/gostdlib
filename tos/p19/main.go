package main

import (
	"fmt"
	"log"
	"os"
)

var lorem = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris 
nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 
reprehenderit in voluptate velit esse cillum dolore eu fugiat 
nulla pariatur. Excepteur sint occaecat cupidatat non proident, 
sunt in culpa qui officia deserunt mollit anim id est laborum.
`

func main() {
	// get cli-input arguments
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Enter a name to create a temp file")
		return
	}
	dirName := args[0]

	f, err := os.CreateTemp(dirName, "tempfile-*-.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	temp := f.Name()

	fmt.Printf("Tempfile name %s\n", temp)

	if _, err := f.Write([]byte(lorem)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
