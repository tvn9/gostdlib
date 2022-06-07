// Opening file by name
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))

	f, err = os.OpenFile("text.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.WriteString(f, "Test string")
}
