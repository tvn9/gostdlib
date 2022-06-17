// CopyN copies n bytes (or until a error) from src to dst. It returns the number
// of bytes copied and the earliest error
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	n, err := io.CopyN(os.Stdout, r, 15)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("Number of byte written:", n)
}
