// Creating a program interface with the flag package
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom type need to implement flag.Value interface to be able
// to use in flag.Var function.
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ArrayValue) Set(str string) error {
	*s = strings.Split(str, ",")
	return nil
}

func main() {
	// Extracting flag values with methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Read the flag using the Var function.
	// In this case the variable must be defined prior to the flag.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through.")

	// Execute the flag.Parse function, to read the flags to defined variables.
	// Without this call the flag variables remain empty.
	flag.Parse()

	// Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
