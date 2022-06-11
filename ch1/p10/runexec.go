// Calling an external process
// The go binary could also be used as a tool for various utilities
// and with use of go run as a replacement for the bash script.
// For these purposes, it is usual that the command-line utilities are called.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println()
		fmt.Println(out.String())
	}
}
