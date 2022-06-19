package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Getwd returns a rooted path name corresponding to the current directory.
	// If the current directory can be reached via multipoe paths (due to symbolic links),
	// Getwd may return any one of them.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("fails to get current directory", err)
	}
	fmt.Printf("Current working directory %s\n", dir)

	// Hostname returns the host name reported by the kernel.
	host, err := os.Hostname()
	if err != nil {
		log.Fatalln("fails to get hostname", err)
	}
	fmt.Printf("hostname %s\n", host)
}
