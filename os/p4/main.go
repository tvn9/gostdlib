package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// read cli-argument
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a file name, then press ENTER")
		return
	}

	file := args[0]

	// verify file existence
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file is not exist", err)
		}
	}
	log.Println("file exist.")

	mtime := time.Date(2020, time.June, 16, 2, 57, 11, 0, time.UTC)
	atime := time.Date(2020, time.May, 15, 9, 35, 12, 0, time.UTC)

	if err := os.Chtimes(file, atime, mtime); err != nil {
		log.Fatal(err)
	}
}
