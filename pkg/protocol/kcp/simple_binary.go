package kcp

import (
	"bytes"
	"encoding/gob"
	"github.com/pmiguel/kiwi/pkg/protocol"
)

func Decode[T protocol.Request | protocol.Response](packet []byte) (T, error) {
	buffer := bytes.NewBuffer(packet)
	dec := gob.NewDecoder(buffer)

	var req T
	err := dec.Decode(&req)

	return req, err
}

func Encode[T protocol.Request | protocol.Response](response *T) ([]byte, error) {
	buffer := bytes.Buffer{}
	dec := gob.NewEncoder(&buffer)

	err := dec.Encode(response)
	return buffer.Bytes(), err
}
