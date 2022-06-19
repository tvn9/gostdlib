package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
Usage: cmdName dirName then press ENTER
the program will read the content of named directory
and list all sub directory and files and  to file2.txt
`

func main() {
	// get cli input
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print(usage)
		return
	}
	dirName := args[0]

	// ReadDir from cli-argument
	// os.ReadDir returns a slice of fs.DirEntry, and error
	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
