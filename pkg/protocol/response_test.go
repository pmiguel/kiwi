package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponse(t *testing.T) {
	expectedContent := "response"

	request := NewResponse(expectedContent, true)

	assert.Equal(t, expectedContent, request.Content, "should be equal")
	assert.True(t, request.Err, "should be true")
}

func TestResponse_MarshalBinary(t *testing.T) {
	expectedContent := "response"

	response := NewResponse(expectedContent, true)

	binaryData, err := response.MarshalBinary()

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "true response\n", string(binaryData), "should be equal")
}

func TestResponse_UnmarshalBinary(t *testing.T) {
	expectedContent := "response"

	input := []byte("true" + " " + expectedContent)

	response := NewResponse(expectedContent, true)

	err := response.UnmarshalBinary(input)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expectedContent, response.Content, "should be equal")
}
