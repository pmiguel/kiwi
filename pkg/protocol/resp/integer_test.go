package resp

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
