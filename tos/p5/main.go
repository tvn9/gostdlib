package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// get executable binary
	str, err := os.Executable()
	if err != nil {
		log.Fatalln("fails to get executable!")
	}
	fmt.Println("Executable is", str)
}
