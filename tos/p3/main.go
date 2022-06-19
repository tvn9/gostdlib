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

	uid := os.Getuid()
	gid := os.Getgid()

	fmt.Println(uid, gid)

	// Change file ownership.
	err = os.Chown("test.txt", uid, gid)
	if err != nil {
		log.Println(err)
	}

}
