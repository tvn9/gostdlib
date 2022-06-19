// Basic file permission using os.Chmod(name string, mode FileMode) error
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
		fmt.Println("Please enter a file [name perm] value, then press ENTER")
		return
	}

	fname := args[0]
	err := os.Chmod(fname, 0766)
	if err != nil {
		log.Fatal(err)
	}
}
