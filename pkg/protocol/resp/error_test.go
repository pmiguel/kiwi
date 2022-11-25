package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Encode_NoSpaces(t *testing.T) {
	input := Error("ERR")
	output := input.Encode()
	assert.Equal(t, "-ERR\r\n", output, "should be equal")
}

func TestError_Encode_WithSpaces(t *testing.T) {
	input := Error("ERR unknown")
	output := input.Encode()
	assert.Equal(t, "-ERR unknown\r\n", output, "should be equal")
}

func TestDecodeError_NoSpaces(t *testing.T) {
	input := "-ERR\r\n"
	output := DecodeError(input)
	assert.Equal(t, Error("ERR"), output, "should be equal")
}

func TestDecodeError_WithSpaces(t *testing.T) {
	input := "-ERR unknown\r\n"
	output := DecodeError(input)
	assert.Equal(t, Error("ERR unknown"), output, "should be equal")
}
