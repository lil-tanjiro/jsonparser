package jsonparser

import (
	"fmt"
	"testing"
)

var (
	data []byte = []byte("\"hi my name is Lorem. i am\\n see you\"")
)

func TestParseString(t *testing.T) {
	var (
		buf  Buffer
		item *JsonItem
	)

	item = new(JsonItem)

	buf.data = data
	buf.lenght = len(data)
	buf.offset = 0

	fmt.Println(ParseString(item, &buf))

	fmt.Println(item)
	fmt.Println(buf)
}
