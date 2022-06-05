package jsonparser

const (
	JsonNull = iota
	JsonTrue
	JsonFalse
	JsonString
)

// var (
// 	nullV  = []byte("null")
// 	trueV  = []byte("true")
// 	falseV = []byte("false")
// )

type JsonItem struct {
	t         int
	stringVar []byte
}

func newJsonItem() *JsonItem {
	return &JsonItem{
		t: 0,
	}
}

type Buffer struct {
	data   []byte
	offset int
	lenght int
}

func (b *Buffer) skipNearSpaces() {
	for b.offset < b.lenght && b.data[b.offset] <= 32 {
		b.offset++
	}
}

func (b *Buffer) canRead(n int) bool {
	return (b.offset + n) <= b.lenght
}

func (b *Buffer) canAccessAtIndex(index int) bool {
	return (b.offset + index) < b.lenght
}

// cases:
// "\n\n\n\"
// ""
// fsdf"
// "fsdf
// "by \"Robert\"\n"
func ParseString(item *JsonItem, buf *Buffer) bool {
	if !buf.canAccessAtIndex(0) {
		return false
	}

	if buf.data[buf.offset] != '"' {
		return false
	}

	var (
		output []byte
		index int
		allocated int
		skiped int
	)

	index = 0
	allocated = 0
	skiped = 0

	for buf.canAccessAtIndex(index) && buf.data[buf.offset + index] != '"' {
		if buf.data[buf.offset+index] == '\\' {
			if 

			index++
			skiped++
		}

		index++
	}

	output = make([]byte, allocated - skiped)

	
}
