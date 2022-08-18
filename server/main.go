package main

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "7170"
	TYPE = "tcp"
)

func handleIncomingRequest(conn net.Conn) {
	fmt.Println("<= " + conn.RemoteAddr().String())
	counter := 0
	for {
		buffer := make([]byte, 1024)
		length, err := conn.Read(buffer)

		if err != nil {
			break
		}

		content := string(buffer[:length])
		fmt.Printf("<< %s (%d bytes) [0x%x], n:%d\n", content, length, content, counter)
		counter++
		conn.Write([]byte("PONG"))
	}
	fmt.Println("=> " + conn.RemoteAddr().String())
	conn.Close()
}

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Listening on port " + PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleIncomingRequest(conn)
	}
}
