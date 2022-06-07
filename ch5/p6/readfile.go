// Reading the file into a string
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	fContent, err := os.ReadFile("temp/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
