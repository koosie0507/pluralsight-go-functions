package readers

import (
	"errors"
	"io"
	"math"
)

type BadReader struct {
	err error
}

func NewBadReader(input string) io.Reader {
	return BadReader{ err: errors.New(input) }
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}


type SimpleReader struct {
	count int
}

func NewSimpleReader(maxCount int) io.Reader {
	return &SimpleReader{count: maxCount}
}

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	if sr.count < 1 {
		return sr.count, io.EOF
	}
	bytes := sr.count - len(p)
	bytesRead := len(p)
	if sr.count < len(p) {
		bytesRead = sr.count
	}
	sr.count = int(math.Max(float64(bytes), 0))
	return bytesRead, nil
}