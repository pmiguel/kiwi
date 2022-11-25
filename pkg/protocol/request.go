package protocol

import (
	"bytes"
	"fmt"
)

type Request struct {
	Command   string
	Arguments []string
	Options   []string
	Key       string
	Value     string
}

func NewRequest(command string, key string, value string) *Request {
	return &Request{Command: command, Key: key, Value: value}
}

func (r *Request) String() string {
	return fmt.Sprintf("%s\n", r.Command)
}

func (r *Request) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanf(b, "%s", &r.Command)
	return err
}

func (r *Request) MarshalBinary() (data []byte, err error) {
	return []byte(r.Command), nil
}
