package main

import (
	"fmt"
	"os"
)

func main() {
	// Getpagesize returns the underlying system's memory page size.
	ps := os.Getpagesize()

	// Getpid returns the process id of the caller
	ppid := os.Getppid()

	fmt.Printf("underlying system's memory page size. %d\n", ps)
	fmt.Printf("process id %d\n", ppid)
}
