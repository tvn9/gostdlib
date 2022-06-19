//WriteFile writes data to the named file, creating it if necessary. If the file does not exist
//WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it
//before writing, without changing permissions.
package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
Usage: cmdName file1.txt file2.txt then press ENTER
the program will read the content of file1.txt and write
to the content to file2.txt
`

func main() {
	// get cli input
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Print(usage)
		return
	}

	fName1, fName2 := args[0], args[1]

	// Reading file content using os.ReadFile
	data, err := os.ReadFile(fName1)
	if err != nil {
		log.Fatal(err)
	}

	// Write file to new file name
	err = os.WriteFile(fName2, []byte(data), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Writing file content to os.Stdout
	os.Stdout.Write(data)
}
