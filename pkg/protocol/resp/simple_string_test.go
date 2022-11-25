package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleString_Encode_NoSpaces(t *testing.T) {
	input := SimpleString("hello")
	output := input.Encode()
	assert.Equal(t, "+hello\r\n", output, "should be equal")
}

func TestDecodeSimpleString_NoSpaces(t *testing.T) {
	input := "+hello\r\n"
	output := DecodeSimpleString(input)
	assert.Equal(t, SimpleString("hello"), output, "should be equal")
}

func TestSimpleString_Encode_WithSpaces(t *testing.T) {
	input := SimpleString("hello world")
	output := input.Encode()
	assert.Equal(t, "+hello world\r\n", output, "should be equal")
}

func TestDecodeSimpleString_WithSpaces(t *testing.T) {
	input := "+hello world\r\n"
	output := DecodeSimpleString(input)
	assert.Equal(t, SimpleString("hello world"), output, "should be equal")
}
