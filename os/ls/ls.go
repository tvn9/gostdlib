package main

import (
	"fmt"
	"log"
	"os"
)

func ListDir(cwd string) {
	cd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to find current directory %s", err)
	}
	fmt.Printf("Current directory: %s\n", cd)
	files, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalf("failed to read directory %s", err)
	}

	for i, file := range files {
		fmt.Printf("%d, %s\n", i, file)
	}
}

func main() {
	ListDir("./")
	// ListDir("../")
}
