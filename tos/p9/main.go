package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Getgroups ids
	// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
	gids, err := os.Getgroups()
	if err != nil {
		log.Fatalln("fails to get group ids", err)
	}

	fmt.Printf("The group's ids that the user belong to %d\n", gids)
}
