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
	fmt.Println("-> Hello " + conn.RemoteAddr().String())
	counter := 0
	for {
		buffer := make([]byte, 1024)
		length, err := conn.Read(buffer)

		if err != nil {
			break
		}

		fmt.Printf("%d << %s [0x%x], len: %d \n", counter, string(buffer[:length]), buffer[:length], length)
		counter++
		conn.Write([]byte("PONG"))
	}
	fmt.Println("-X Bye " + conn.RemoteAddr().String())
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
