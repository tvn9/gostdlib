package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the user group id
	egid := os.Getegid()
	euid := os.Geteuid()

	// Get the user env
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("Numberic effiective group id %d\n", egid)
	fmt.Printf("Numberic effiective user id %d\n", euid)
	fmt.Printf("Environment variable named by the key %q %q\n", os.Getenv("NAME"), os.Getenv("BURROW"))
}
