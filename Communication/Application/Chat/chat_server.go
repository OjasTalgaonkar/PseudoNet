package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

func WriteMessage(msg string) []byte {
	length := uint32(len(msg))
	buf := make([]byte, 4+len(msg))

	// Put the 4-byte length in front (big endian)
	binary.BigEndian.PutUint32(buf[:4], length)

	copy(buf[4:], msg)

	return buf
}

func acceptClients(ln net.Listener) {

	var (
		clients      = make(map[net.Conn]bool)
		clientsMutex = sync.Mutex{} //prevent data races
	)

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("error : ", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		clientsMutex.Lock()
		clients[conn] = true
		clientsMutex.Unlock()
	}

}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Connection error, verify your connection status and rerun")
	}
	defer ln.Close()

}
