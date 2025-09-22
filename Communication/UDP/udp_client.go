package main

import (
	"fmt"
	"net"
)

func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	for {
		_, err = conn.Write([]byte("ping"))
		if err != nil {
			fmt.Println("Error sending:", err)
			return
		}
		fmt.Println("Sent: ping")

		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Println("Received:", string(buf[:n]))
	}
}
