package binarystdout

import (
	"github.com/pkg/errors"

	"io"
	"os"
	"sync"
)

var binarystdout *BinaryStdOut
var once sync.Once

// BinaryStdOut provides methods for converting primtive type variables (bool, byte, int16, int, int64,string)
// to sequences of bits and writing them to standard output. Uses big-endian (most-significant byte first).
type BinaryStdOut struct {
	out    io.Writer // output stream (standard output)
	buffer int       // 8-bit buffer of bits to write
	n      int       // number of bits remaining in buffer
}

// NewBinaryStdOut creates the singleton BinaryStdout struct
func NewBinaryStdOut() *BinaryStdOut {
	once.Do(func() {
		binarystdout = &BinaryStdOut{out: os.Stdout, buffer: 0, n: 0}
	})
	return binarystdout
}

// WriteBit writes the specified bit to standard output.
func (bs *BinaryStdOut) WriteBit(bit bool) {
	bs.buffer <<= 1
	if bit {
		bs.buffer |= 1
	}
	bs.n++
	if bs.n == 8 {
		bs.clearBuffer()
	}
}

// WriteBitR writes the r-bit int to standard output.
func (bs *BinaryStdOut) WriteBitR(x int, r int) error {
	if r == 32 {
		err := bs.WriteInt(x)
		if err != nil {
			return err
		}
		return nil
	}
	if r < 1 || r > 32 {
		return errors.Errorf("illegal value for r = %d\n", r)
	}
	if x < 0 || x >= (1<<r) {
		return errors.Errorf("illegal value for x = %d\n", x)
	}
	for i := 1; i <= r; i++ {
		bit := ((x >> (r - i)) & 1) == 1
		bs.WriteBit(bit)
	}
	return nil
}

// WriterInt writes the 32-bit int to standard output.
func (bs *BinaryStdOut) WriteInt(i int) error {
	x := uint(i)
	err1 := bs.WriteByte(byte((x >> 24) & 0xff))
	err2 := bs.WriteByte(byte((x >> 16) & 0xff))
	err3 := bs.WriteByte(byte((x >> 8) & 0xff))
	err4 := bs.WriteByte(byte((x >> 0) & 0xff))
	return checkError(err1, err2, err3, err4)
}

// WriteInt16 Writes the 16-bit int to standard output.
func (bs *BinaryStdOut) WriteInt16(i int16) error {
	x := uint16(i)
	err1 := bs.WriteByte(byte((x >> 8) & 0xff))
	err2 := bs.WriteByte(byte((x >> 0) & 0xff))
	return checkError(err1, err2)
}

// WriteInt64 writes the 64-bit int to standard output.
func (bs *BinaryStdOut) WriteInt64(i int64) error {
	x := uint64(i)
	err1 := bs.WriteByte(byte((x >> 56) & 0xff))
	err2 := bs.WriteByte(byte((x >> 48) & 0xff))
	err3 := bs.WriteByte(byte((x >> 40) & 0xff))
	err4 := bs.WriteByte(byte((x >> 32) & 0xff))
	err5 := bs.WriteByte(byte((x >> 24) & 0xff))
	err6 := bs.WriteByte(byte((x >> 16) & 0xff))
	err7 := bs.WriteByte(byte((x >> 8) & 0xff))
	err8 := bs.WriteByte(byte((x >> 0) & 0xff))
	return checkError(err1, err2, err3, err4, err5, err6, err7, err8)
}

// WriteByte writes the 8-bit byte to standard output.
func (bs *BinaryStdOut) WriteByte(x byte) error {
	if bs.n == 0 {
		_, err := bs.out.Write([]byte{x})
		if err != nil {
			return err
		}
		return nil
	}

	for i := 1; i <= 8; i++ {
		bit := ((int(x) >> (8 - i)) & 1) == 1
		bs.WriteBit(bit)
	}
	return nil
}

// WriteString writes the string of 8-bit characters to standard output.
func (bs *BinaryStdOut) WriteString(s string) error {
	for i := range s {
		if err := bs.WriteByte(s[i]); err != nil {
			return err
		}
	}
	return nil
}

// Close flushes and closes standard output.
func (bs *BinaryStdOut) Close() {
	bs.clearBuffer()
}

func (bs *BinaryStdOut) clearBuffer() {
	if bs.n == 0 {
		return
	}
	if bs.n > 0 {
		bs.buffer <<= (8 - bs.n)
	}

	_, err := bs.out.Write([]byte{byte(bs.buffer)})
	if err != nil {
		panic(err)
	}
	bs.n = 0
	bs.buffer = 0
}

func checkError(errs ...error) error {
	for i := range errs {
		if errs[i] != nil {
			return errs[i]
		}
	}
	return nil
}
