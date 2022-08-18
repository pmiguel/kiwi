package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pmiguel/kiwi/common/protocol"
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

	var writeBuffer bytes.Buffer
	enc := gob.NewEncoder(&writeBuffer)

	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		req := protocol.NewRequest(strings.TrimSpace(text), "", "")
		enc.Encode(req)
		conn.Write(writeBuffer.Bytes())

		buffer := make([]byte, 1024)
		length, _ := conn.Read(buffer)
		fmt.Printf("<< %s \n", string(buffer[:length]))
	}
	conn.Close()
}
