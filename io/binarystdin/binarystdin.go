package binarystdin

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

const eof = -1

var binaryStdIn *BinaryStdIn
var once sync.Once

type BinaryStdIn struct {
	in            io.Reader
	buffer        int
	n             int
	isInitialized bool
}

func NewBinaryStdIn() *BinaryStdIn {
	once.Do(func() {
		binaryStdIn = &BinaryStdIn{
			in:            os.Stdin,
			buffer:        0,
			n:             0,
			isInitialized: false,
		}
	})
	return binaryStdIn
}

func (bs *BinaryStdIn) initialize() {
	bs.fillBuffer()
	bs.isInitialized = true
}

func (bs *BinaryStdIn) fillBuffer() {
	p := make([]byte, 1)
	if _, err := bs.in.Read(p); err != nil {
		if err == io.EOF {
			bs.buffer = eof
			bs.n = -1
		} else {
			panic(err)
		}
	} else {
		bs.n = 8
		bs.buffer = int(p[0])
	}
}

func (bs *BinaryStdIn) IsEmpty() bool {
	if !bs.isInitialized {
		bs.initialize()
	}
	return bs.buffer == eof
}

func (bs *BinaryStdIn) ReadBool() (bool, error) {
	if bs.IsEmpty() {
		return false, errors.New("reading from empty input stream")

	}
	bs.n--
	bit := ((bs.buffer >> bs.n) & 1) == 1
	if bs.n == 0 {
		bs.fillBuffer()
	}
	return bit, nil
}

func (bs *BinaryStdIn) ReadByte() (byte, error) {
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")
	}
	if bs.n == 8 {
		x := byte(bs.buffer)
		bs.fillBuffer()
		return x, nil
	}
	x := bs.buffer
	x <<= (8 - bs.n)
	oldN := bs.n
	bs.fillBuffer()
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")

	}
	bs.n = oldN
	x |= (bs.buffer >> bs.n) // ?????
	return byte(x), nil
}

func (bs *BinaryStdIn) ReadInt() (int, error) {
	x := 0
	for i := 0; i < 4; i++ {
		b, err := bs.ReadByte()
		if err != nil {
			return -1, err
		}
		x <<= 8
		x |= int(b)
	}
	return x, nil
}

func (bs *BinaryStdIn) ReadIntR(r int) (int, error) {
	if r < 1 || r > 32 {
		return -1, errors.Errorf("illegal value of r = %d\n", r)
	}
	if r == 32 {
		return bs.ReadInt()
	}
	x := 0
	for i := 0; i < r; i++ {
		x <<= 1
		bit, err := bs.ReadBool()
		if err != nil {
			return -1, err
		}
		if bit {
			x |= 1
		}
	}
	return x, nil
}

func (bs *BinaryStdIn) ReadString() (string, error) {
	if bs.IsEmpty() {
		return "", errors.New("reading from empty input stream")
	}
	var sb strings.Builder
	for !bs.IsEmpty() {
		b, err := bs.ReadByte()
		if err != nil {
			return "", err
		}
		sb.WriteByte(b)
	}
	return sb.String(), nil
}
