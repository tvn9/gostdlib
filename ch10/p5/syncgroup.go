// Synchronizing goroutines with WaitGroup
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Group 1, Exiting %d\n", idx)
		}(i)
	}
	// wg.Wait()
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Group 2, Exiting %d\n", idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done.")
}
