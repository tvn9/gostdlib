// Resolving local IP addresses
package main

import (
	"fmt"
	"net"
)

func main() {

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		addr, err := interf.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println("interface", interf.Name)
		for _, add := range addr {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}

	ifAddrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for i, addr := range ifAddrs {
		fmt.Printf("Net %d, addr: %s\n", i, addr)
	}
}
