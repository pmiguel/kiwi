package protocol

import "testing"

func TestNewRequest(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	request := NewRequest(expectedCommand, expectedKey, expectedValue)

	if request.Command != expectedCommand {
		t.Errorf("request.Command = %s; want %s", request.Command, expectedCommand)
	}

	if request.Key != expectedKey {
		t.Errorf("request.Key = %s; want %s", request.Key, expectedKey)
	}

	if request.Value != expectedValue {
		t.Errorf("request.Value = %s; want %s", request.Value, expectedValue)
	}
}

func TestRequest_MarshalBinary(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	request := NewRequest(expectedCommand, expectedKey, expectedValue)

	binaryData, err := request.MarshalBinary()

	if err != nil {
		t.Errorf("Error. %s", err)
	}

	if string(binaryData) != "command key value\n" {
		t.Errorf("Failed serialization failed. Got %s", string(binaryData))
	}
}
func TestRequest_UnmarshalBinary(t *testing.T) {
	expectedCommand := "command"
	expectedKey := "key"
	expectedValue := "value"

	input := []byte(expectedCommand + " " + expectedKey + " " + expectedValue)

	request := NewRequest("", "", "")

	err := request.UnmarshalBinary(input)

	if err != nil {
		t.Errorf("Error. %s", err)
	}

	if request.Command != expectedCommand {
		t.Errorf("request.Command = %s; want %s", request.Command, expectedCommand)
	}

	if request.Key != expectedKey {
		t.Errorf("request.Key = %s; want %s", request.Key, expectedKey)
	}

	if request.Value != expectedValue {
		t.Errorf("request.Value = %s; want %s", request.Value, expectedValue)
	}
}
