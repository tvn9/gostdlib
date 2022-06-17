package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	// validates user input
	if len(args) != 1 {
		fmt.Println("Please a name to create a directory")
		return
	}
	dirName := args[0]

	// Mkdir make a directory with access permission mode
	err := os.Mkdir(dirName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile(dirName+"/testfile.txt", []byte("Hello, Gophers!\n"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
