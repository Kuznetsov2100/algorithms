package binarystdout

import (
	"errors"
	"io"
	"os"
	"sync"
)

var binarystdout *BinaryStdOut
var once sync.Once

type BinaryStdOut struct {
	out    io.Writer
	buffer int
	n      int
}

func NewBinaryStdOut() *BinaryStdOut {
	once.Do(func() {
		binarystdout = &BinaryStdOut{
			out:    os.Stdout,
			buffer: 0,
			n:      0,
		}
	})
	return binarystdout
}

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

func (bs *BinaryStdOut) WriteBits(x int, r int) error {
	if r == 32 {
		err := bs.WriteInt(x)
		if err != nil {
			return err
		}
		return nil
	}
	if r < 1 || r > 32 {
		return errors.New("invalid r")
	}
	if x < 0 || x >= (1<<r) {
		return errors.New("invalid x")
	}
	for i := 0; i < r; i++ {
		bit := ((x >> (r - i - 1)) & 1) == 1
		bs.WriteBit(bit)
	}
	return nil
}

func (bs *BinaryStdOut) WriteInt(x int) error {
	err := bs.WriteByte(byte((x >> 24) & 0xff))
	if err != nil {
		return err
	}
	err = bs.WriteByte(byte((x >> 16) & 0xff))
	if err != nil {
		return err
	}
	err = bs.WriteByte(byte((x >> 8) & 0xff))
	if err != nil {
		return err
	}
	err = bs.WriteByte(byte((x >> 0) & 0xff))
	if err != nil {
		return err
	}
	return nil

}

func (bs *BinaryStdOut) WriteByte(x byte) error {
	if bs.n == 0 {
		_, err := bs.out.Write([]byte{x})
		if err != nil {
			return err
		}
		return nil
	}

	for i := 0; i < 8; i++ {
		bit := ((int(x) >> (8 - i - 1)) & 1) == 1
		bs.WriteBit(bit)
	}
	return nil
}

func (bs *BinaryStdOut) WriteString(s string) error {
	for i := range s {
		if err := bs.WriteByte(s[i]); err != nil {
			return err
		}
	}
	return nil
}

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
