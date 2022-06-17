// MkdirAll creates a directory named path, along with any necessary
// parents and returns nil, or else returns an error. The permission bits
// perm (before umask) are used for all directories that MkdirAll creates.
// if path is already a directory, MkdirAll does nothing and return nil.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Get command-line input
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a name to create directory")
		fmt.Println("Example: mkdirs dir1 will create dir1 in current directory")
		fmt.Println("mkdirs dir1/dir2/dir3 will create a dir1 and sub dir2/dir3")
		return
	}
	dirName := args[0]

	err := os.MkdirAll(dirName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = os.WriteFile(dirName+"/testfile.txt", []byte("Hello, Gophers!\n"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
