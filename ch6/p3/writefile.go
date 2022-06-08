// Writing files
package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is awesome! ")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Yeah!, Go is greate."))
	if err != nil {
		panic(err)
	}
}
