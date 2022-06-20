package main

import "log"

func main() {
	c := make(chan int, 3)
	c <- 1

	log.Println(<-c) // prints 1

	c <- 2
	close(c)

	log.Println(<-c) // Prints 2
	log.Println(<-c) // Prints 0

	if i, ok := <-c; ok {
		log.Printf("Channel is open, got %d", i)
	} else {
		log.Printf("Channel is closed, go %d", i)
	}

	close(c) // Panics, channel is already closed
}
