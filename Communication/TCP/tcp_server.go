package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Connection error, verify your connection status and rerun")
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("error : ", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("error : ", err)
			break
		}

		fmt.Println("Message : ", message)

		conn.Write([]byte(message))
	}
}
