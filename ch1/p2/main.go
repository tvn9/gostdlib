package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// This call will print all command arguments.
	fmt.Println(args)

	// The first arguments index 0 from the slice is the name of the called binary.
	programName := args[0]
	fmt.Printf("The binary is %s:\n", programName)

	if len(args) < 2 {
		fmt.Println("Try to enter some words after the command to see more info.")
		return
	}

	// The rest of the command arguments can be obtain by [1:]
	cmdArgs := args[1:]
	fmt.Printf("The rest of the command arguments %s\n", cmdArgs)

	// Just to test the for loop on the args slice
	for i, a := range args {
		fmt.Printf("%d item %s\n", i, a)
	}
}
