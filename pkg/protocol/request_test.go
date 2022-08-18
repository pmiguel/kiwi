package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRequest(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	request := NewRequest(expectedCommand, expectedKey, expectedValue)

	assert.Equal(t, expectedCommand, request.Command, "should be equal")
	assert.Equal(t, expectedKey, request.Key, "should be equal")
	assert.Equal(t, expectedValue, request.Value, "should be equal")
}

func TestRequest_MarshalBinary(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	request := NewRequest(expectedCommand, expectedKey, expectedValue)

	binaryData, err := request.MarshalBinary()

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "command key value\n", string(binaryData), "should be equal")
}

func TestRequest_UnmarshalBinary(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	input := []byte(expectedCommand + " " + expectedKey + " " + expectedValue)

	request := NewRequest("", "", "")

	err := request.UnmarshalBinary(input)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expectedCommand, request.Command, "should be equal")
	assert.Equal(t, expectedKey, request.Key, "should be equal")
	assert.Equal(t, expectedValue, request.Value, "should be equal")
}
