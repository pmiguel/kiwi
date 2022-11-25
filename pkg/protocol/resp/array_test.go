package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Encode_Strings(t *testing.T) {
	var input = Array{BulkString("hello"), BulkString("world")}
	output := input.Encode()

	assert.Equal(t, "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", output, "should be equal")
}

func TestDecodeArray(t *testing.T) {
	input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	output := DecodeArray(input)
	expected := []string{"hello", "world"}
	assert.Equal(t, expected, output)
}
