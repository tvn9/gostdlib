// Creating the UDP server
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	buffer := make([]byte, 2048)
	fmt.Println("Waiting for client...")
	for {
		_, addr, err := pc.ReadFrom(buffer)
		if err == nil {
			rcvMsg := string(buffer)
			fmt.Println("Received: " + rcvMsg)
			if _, err := pc.WriteTo([]byte("Received: "+rcvMsg), addr); err != nil {
				fmt.Println("error on write: " + err.Error())
			}
		} else {
			fmt.Println("error: " + err.Error())
		}
	}
}
