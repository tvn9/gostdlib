// Getting and setting environment variables with default values
// This example demonstrate how to read, set and unset the environment.
// it also show you how to implement the default option if the variables
// is not set.
package main

import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}
