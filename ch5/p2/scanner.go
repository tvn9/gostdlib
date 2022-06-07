// Read input by line using bufio.NewScanner
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// The scanner is able to scan input by lines
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
