package encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt_Encode_PositiveNumber(t *testing.T) {
	input := Int(123)
	output := input.Encode()
	assert.Equal(t, ":123\r\n", output, "should be equal")
}

func TestInt_Encode_NegativeNumber(t *testing.T) {
	input := Int(-123)
	output := input.Encode()
	assert.Equal(t, ":-123\r\n", output, "should be equal")
}

func TestInt_Decode_PositiveNumber(t *testing.T) {
	input := ":123\r\n"
	output := DecodeInt(input)
	assert.Equal(t, Int(123), output, "should be equal")
}

func TestInt_Decode_NegativeNumber(t *testing.T) {
	input := ":-123\r\n"
	output := DecodeInt(input)
	assert.Equal(t, Int(-123), output, "should be equal")
}

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

func TestArray_Encode_Strings(t *testing.T) {
	var input = Array{BulkString("hello"), BulkString("world")}
	output := input.Encode()

	assert.Equal(t, "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", output, "should be equal")
}

func TestDecodeLength_PositiveNumber(t *testing.T) {
	input := "12\r\n"
	output, index := DecodeLength([]rune(input), 0)

	assert.Equal(t, 12, output)
	assert.Equal(t, 4, index)
}

func TestDecodeLength_NegativeNumber(t *testing.T) {
	input := "-12\r\n"
	output, index := DecodeLength([]rune(input), 0)

	assert.Equal(t, -12, output)
	assert.Equal(t, 5, index)
}

func TestDecodeArray(t *testing.T) {
	input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	output := DecodeArray(input)
	expected := []string{"hello", "world"}
	assert.Equal(t, expected, output)
}
