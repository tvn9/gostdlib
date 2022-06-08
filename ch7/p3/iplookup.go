// Resolving the domain by IP address and vice versa
package main

import (
	"fmt"
	"net"
)

func main() {
	// Resole by IP address
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// Resole by localhost
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
