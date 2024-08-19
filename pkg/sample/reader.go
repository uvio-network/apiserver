package sample

import (
	"fmt"
	"io"
)

type fake struct {
	byt []byte
	pos int
}

func newFakeReader(num int) io.Reader {
	return &fake{byt: []byte(fmt.Sprintf("%d", num))}
}

func (f *fake) Read(p []byte) (n int, err error) {
	if f.pos >= len(f.byt) {
		return 0, io.EOF
	}

	n = copy(p, f.byt[f.pos:])
	f.pos += n

	return n, nil
}
