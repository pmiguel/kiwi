package resp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
