package resp

import "fmt"

const prefixArray = '*'

type Array []Resp

func (arr *Array) Encode() string {
	if arr == nil {
		return fmt.Sprintf("%c%d%s", prefixArray, -1, sepMarker)
	}

	resolved := *arr

	itemCount := len(resolved)
	buff := fmt.Sprintf("%c%d%s", prefixArray, itemCount, sepMarker)
	for i := 0; i < itemCount; i++ {
		value := resolved[i].Encode()
		buff += fmt.Sprintf("%s", value)
	}
	return buff
}

func DecodeArray(input string) []string {
	runes := []rune(input)

	if runes[0] != prefixArray {
		return nil
	}

	itemCount, currentIndex := DecodeLength(runes, 1)

	dataBuffer := make([]string, itemCount)

	for i := 0; i < itemCount; i++ {
		itemType := runes[currentIndex]
		switch itemType {
		case prefixBulkString:
			stringLength, index := DecodeLength(runes, currentIndex+1)
			buff := make([]rune, stringLength)
			for j := 0; j < stringLength; j++ {
				buff[j] = runes[index+j]
			}
			dataBuffer[i] = string(buff)
			currentIndex = index + stringLength + 2
		}
	}
	return dataBuffer
}
