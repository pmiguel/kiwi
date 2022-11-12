package encoding

import (
	"fmt"
)

type ContentType uint8

const (
	SimpleString ContentType = iota
	Error
	Integer
	BulkString
)

const (
	prefixSimpleString = '+'
	prefixError        = '-'
	prefixInteger      = ':'
	prefixBulkString   = '$'
	prefixArray        = '*'
	separator          = "\r\n"
)

type Primitive struct {
	PrimitiveType ContentType
	Content       *string
}

func Int64ToPrimitive(input int64) Primitive {
	str := fmt.Sprint(input)
	return Primitive{PrimitiveType: Integer, Content: &str}
}

func StringToPrimitive(input string) Primitive {
	str := fmt.Sprint(input)
	return Primitive{PrimitiveType: SimpleString, Content: &str}
}

func (p *Primitive) Encode() string {
	if p.Content != nil {
		switch p.PrimitiveType {
		case SimpleString:
			return fmt.Sprintf("%c%s%s", prefixSimpleString, *p.Content, separator)
		case Error:
			return fmt.Sprintf("%c%s%s", prefixError, *p.Content, separator)
		case Integer:
			return fmt.Sprintf("%c%s%s", prefixInteger, *p.Content, separator)
		case BulkString:
			return fmt.Sprintf("%c%d%s%s%s", prefixBulkString, len(*p.Content), separator, *p.Content, separator)
		}
	}

	return fmt.Sprintf("%c%d%s", prefixBulkString, -1, separator)
}

type ContentArray struct {
	Content []Primitive
}

func (c *ContentArray) Encode() string {
	if c.Content == nil {
		return fmt.Sprintf("%c%d%s", prefixArray, -1, separator)
	}
	itemCount := len(c.Content)
	buff := fmt.Sprintf("%c%d%s", prefixArray, itemCount, separator)
	for i := 0; i < itemCount; i++ {
		value := c.Content[i].Encode()
		buff += fmt.Sprintf("%s", value)
	}
	return buff
}
