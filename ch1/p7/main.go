// Retrieving the current working directory
// Another useful resource of information for the application is the directory,
// where the program binary is located. With this information, the program can
// access the assets and files collocated with the binary file.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to executable file
	fmt.Println("Executable binary path:", ex)

	// Resolve the directory of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :", exPath)

	// Use EvalSymlinks to get the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:", realPath)
}
