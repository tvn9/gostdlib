// Use a Scanner with a custom split function (built by wrapping ScanWords) to validate 32-bit decimal input.
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "1234 5678 12345678901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		//fmt.Println(token)
		//fmt.Println(advance)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
	}
}
