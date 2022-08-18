package protocol

import (
	"bytes"
	"encoding/gob"
)

func Decode[T Request | Response](packet []byte) (T, error) {
	buffer := bytes.NewBuffer(packet)
	dec := gob.NewDecoder(buffer)

	var req T
	err := dec.Decode(&req)

	return req, err
}

func Encode[T Request | Response](response *T) ([]byte, error) {
	buffer := bytes.Buffer{}
	dec := gob.NewEncoder(&buffer)

	err := dec.Encode(response)
	return buffer.Bytes(), err
}
