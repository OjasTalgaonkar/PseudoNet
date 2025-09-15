package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close()

	for {

		conn.Write([]byte("Hello\n"))

		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Server closed connection:", err)
			return
		}
		fmt.Println("Server says:", string(response[:n]))

	}

}
