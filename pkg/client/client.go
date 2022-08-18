package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"log"
	"net"
	"os"
	"strings"
	"time"
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
		var writeBuffer bytes.Buffer
		enc := gob.NewEncoder(&writeBuffer)

		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		target := strings.TrimSpace(text)

		req := protocol.NewRequest(target, "key", "value")
		enc.Encode(req)
		conn.Write(writeBuffer.Bytes())

		var buffer bytes.Buffer
		conn.Read(buffer.Bytes())
		fmt.Printf("<< %s \n", buffer.String())
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}
