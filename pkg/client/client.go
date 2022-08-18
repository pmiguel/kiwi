package main

import (
	"bufio"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "7170"
	TYPE = "tcp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial(TYPE, HOST+":"+PORT)

	if err != nil {
		log.Fatal(err)
	}

	for {
		var readBuffer = make([]byte, 1024)

		fmt.Print(conn.RemoteAddr().String() + " >> ")
		text, _ := reader.ReadString('\n')

		target := strings.TrimSpace(text)

		requestBytes, _ := protocol.Encode[protocol.Request](protocol.NewRequest(target, "key", "value"))
		conn.Write(requestBytes)
		conn.Read(readBuffer)
		dec, _ := protocol.Decode[protocol.Response](readBuffer)

		fmt.Printf("%s (Err: %t)\n", dec.Content, dec.Err)
	}
	conn.Close()
}
