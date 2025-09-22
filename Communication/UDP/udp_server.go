package main

import (
	"fmt"
	"net"
)

func main() {

	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}

	ln, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	defer ln.Close()

	buffer := make([]byte, 1024)
	for {
		size, addr, err := ln.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error : ", err)
		}
		fmt.Println("Recieved ", size, " bytes from ", addr.String(), " : ", string(buffer[:size]))

		_, err = ln.WriteToUDP([]byte("pong"), addr)
		if err != nil {
			fmt.Println("Error sending reply:", err)
			return
		}
	}

}
