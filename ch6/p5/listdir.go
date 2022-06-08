// Listing a directory
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("List by readDir")
	listDirByReadDir(".")
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(".")
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		// Walk the given dir without printing out.
		if wPath == path {
			return nil
		}

		// If given path is folder
		// stop list recursively and print as folder.
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		// Print file name
		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

func listDirByReadDir(path string) {
	lst, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, v := range lst {
		if v.IsDir() {
			fmt.Printf("[%s]\n", v.Name())
		} else {
			fmt.Println(v.Name())
		}
	}
}
