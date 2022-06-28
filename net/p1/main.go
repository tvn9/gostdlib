package main

import (
	"fmt"
	"net"
)

func main() {
	// Get all network interfaces on the system
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Resolve addresses for each interface
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Interface name: %s\n", interf.Name)
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok {
				fmt.Printf("\tIP Addr: %v\n", ip)
			}
		}
	}
}
