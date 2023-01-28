package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 0; port < 100; port++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))

		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("Port %d is open", port)
	}
}
