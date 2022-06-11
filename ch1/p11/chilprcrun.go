// Retrieving child process information
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	proc.Wait()

	// No process state is returned til the process finished.
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState.Pid())

	// The PID could be obtain event for the running process
	fmt.Printf("PID of running process: %d\n\n", proc.ProcessState.Pid())
}
