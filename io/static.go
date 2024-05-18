package io

import (
	"errors"
	"io"
)

type Buffer struct {
	eof   bool
	bytes []byte
}

func NewBuffer(s string) *Buffer {
	buf := new(Buffer)
	buf.bytes = []byte(s)
	return buf
}

func (buf *Buffer) Read(p []byte) (int, error) {
	if p == nil {
		return 0, errors.New("invalid argument: input []byte is nil")
	}
	if buf.eof {
		return 0, io.EOF
	}
	return copy(p, buf.bytes), nil
}
