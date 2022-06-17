// Pipe
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Println("Something went wrong", err)
	}
	fmt.Println("Just wanted to see what coming out:", n)
}
