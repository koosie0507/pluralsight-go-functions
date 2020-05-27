package readers

import (
	"errors"
	"fmt"
	"io"
	"math"
)

type BadReader struct {
	err error
}

func NewBadReader(input string) io.ReadCloser {
	return BadReader{ err: errors.New(input) }
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func (br BadReader) Close() error {
	fmt.Println("Closed bad reader")
	return nil
}


type SimpleReader struct {
	count int
}

func NewSimpleReader(maxCount int) io.ReadCloser {
	return &SimpleReader{count: maxCount}
}

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	if sr.count < 1 {
		panic("reading too much")
	}
	bytes := sr.count - len(p)
	bytesRead := len(p)
	if sr.count < len(p) {
		bytesRead = sr.count
	}
	sr.count = int(math.Max(float64(bytes), 0))
	return bytesRead, nil
}

func (sr *SimpleReader) Close() error {
	fmt.Println("Closed simple reader")
	return nil
}

