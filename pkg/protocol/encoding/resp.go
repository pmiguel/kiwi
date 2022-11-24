package encoding

import (
	"fmt"
)

type Int int64
type SimpleString string
type BulkString string
type Error string
type Array []Resp

type Resp interface {
	Encode() string
}

const (
	prefixSimpleString = '+'
	prefixError        = '-'
	prefixInteger      = ':'
	prefixBulkString   = '$'
	prefixArray        = '*'
	separator          = "\r\n"
)

func (i Int) Encode() string {
	return fmt.Sprintf("%c%d%s", prefixInteger, i, separator)
}

func DecodeInt(input string) Int {
	var out Int
	_, err := fmt.Sscanf(input, ":%d\r\n", &out)
	if err != nil {
		return Int(0)
	}
	return Int(out)
}

func (str SimpleString) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixSimpleString, str, separator)
}

func DecodeSimpleString(input string) SimpleString {
	chars := []rune(input)

	if chars[0] != prefixSimpleString {
		return ""
	}

	buf := make([]rune, len(chars)-3)
	for i := 1; chars[i] != '\r' && i < len(chars)-1; i++ {
		buf[i-1] = chars[i]
	}

	return SimpleString(buf)
}

func (str BulkString) Encode() string {
	return fmt.Sprintf("%c%d%s%s%s", prefixBulkString, len(str), separator, str, separator)
}

func DecodeBulkString(input string) BulkString {
	chars := []rune(input)

	if chars[0] != prefixBulkString {
		return ""
	}

	length, index := DecodeLength(chars, 1)
	buff := make([]rune, length)
	for i := 0; i < length; i++ {
		buff[i] = chars[index+i]
	}
	return BulkString(buff)
}

func (str Error) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixError, str, separator)
}

func DecodeError(input string) Error {
	chars := []rune(input)

	if chars[0] != prefixError {
		return ""
	}

	buf := make([]rune, len(chars)-3)
	for i := 1; chars[i] != '\r' && i < len(chars)-1; i++ {
		buf[i-1] = chars[i]
	}
	return Error(buf)
}

func (arr *Array) Encode() string {
	if arr == nil {
		return fmt.Sprintf("%c%d%s", prefixArray, -1, separator)
	}

	resolved := *arr

	itemCount := len(resolved)
	buff := fmt.Sprintf("%c%d%s", prefixArray, itemCount, separator)
	for i := 0; i < itemCount; i++ {
		value := resolved[i].Encode()
		buff += fmt.Sprintf("%s", value)
	}
	return buff
}

// * 2 \r\n $5\r\nhello\r\n$5\r\nworld\r\n

func DecodeLength(input []rune, startIndex int) (length int, currentIndex int) {
	length = 0
	signal := 1
	currentIndex = startIndex
	for input[currentIndex] != '\r' {
		if currentIndex == 0 && input[currentIndex] == '-' {
			signal = -(signal)
		} else {
			length = (length * 10) + (int(input[currentIndex]) - int('0'))
		}
		currentIndex++
	}
	return length * signal, currentIndex + 2 // Skip \r\n
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
