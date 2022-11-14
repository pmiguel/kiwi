package encoding

import (
	"fmt"
)

type Int int64
type SimpleString string
type BulkString string
type Error string
type Array []RestType

type RestType interface {
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

func (i *Int) Encode() string {
	return fmt.Sprintf("%c%d%s", prefixInteger, i, separator)
}

func DecodeInt(input string) Int {
	var out Int
	_, err := fmt.Sscanf(input, ":%d\r\n", &out)
	if err != nil {
		return 0
	}
	return out
}

func (str *SimpleString) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixSimpleString, *str, separator)
}

func DecodeSimpleString(input string) SimpleString {
	var out SimpleString
	_, err := fmt.Sscanf(input, "+%s\r\n", &out)
	if err != nil {
		return ""
	}
	return out
}

func (str *BulkString) Encode() string {
	return fmt.Sprintf("%c%d%s%s%s", prefixBulkString, len(*str), separator, *str, separator)
}

func DecodeBulkString(input string) BulkString {
	var out BulkString
	var length int
	_, err := fmt.Sscanf(input, "$%d\r\n%s\r\n", &length, &out)
	if err != nil {
		return ""
	}
	return out
}

func (str *Error) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixError, *str, separator)
}

func DecodeError(input string) Error {
	var out Error
	_, err := fmt.Sscanf(input, "+%s\r\n", &out)
	if err != nil {
		return ""
	}
	return out
}

func (arr Array) Encode() string {
	if arr == nil {
		return fmt.Sprintf("%c%d%s", prefixArray, -1, separator)
	}
	itemCount := len(arr)
	buff := fmt.Sprintf("%c%d%s", prefixArray, itemCount, separator)
	for i := 0; i < itemCount; i++ {
		value := arr[i].Encode()
		buff += fmt.Sprintf("%s", value)
	}
	return buff
}
