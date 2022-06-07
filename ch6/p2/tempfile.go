// Creating temporaly files
package main

import (
	"fmt"
	"os"
)

func main() {
	tFile, err := os.CreateTemp("", "gostdcookbook")
	if err != nil {
		panic(err)
	}
	// The call is sponsible for hangling the clean up.
	defer os.Remove(tFile.Name())
	fmt.Println(tFile.Name())
	// TempDir returns
	// the path in string.
	tDir, err := os.MkdirTemp("", "gostdcookbook")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)
	fmt.Println(tDir)
}
