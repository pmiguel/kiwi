package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBulkString_Encode_NoSpaces(t *testing.T) {
	input := BulkString("hello")
	output := input.Encode()
	assert.Equal(t, "$5\r\nhello\r\n", output, "should be equal")
}

func TestDecodeBulkString_NoSpaces(t *testing.T) {
	input := "$5\r\nhello\r\n"
	output := DecodeBulkString(input)
	assert.Equal(t, BulkString("hello"), output, "should be equal")
}

func TestBulkString_Encode_WithSpaces(t *testing.T) {
	input := BulkString("hello world")
	output := input.Encode()
	assert.Equal(t, "$11\r\nhello world\r\n", output, "should be equal")
}

func TestDecodeBulkString_WithSpaces(t *testing.T) {
	input := "$11\r\nhello world\r\n"
	output := DecodeBulkString(input)
	assert.Equal(t, BulkString("hello world"), output, "should be equal")
}
