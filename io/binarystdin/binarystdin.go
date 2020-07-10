package binarystdin

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var binaryStdIn *BinaryStdIn
var once sync.Once

// BinaryStdIn struct provides methods for reading in bits from standard input,
// either one bit at a time (as a bool), 8 bits at a time (as a byte), 16 bits at a time (as a int16),
// 32 bits at a time (as an int), or 64 bits at a time (as a int64).
// All primitive types are assumed to be represented using their golang standard representations,
// in big-endian (most significant byte first) order.
type BinaryStdIn struct {
	in            io.Reader // input stream
	buffer        int       // one byte buffer
	n             int       // number of bits left in buffer
	isInitialized bool      // has BinaryStdIn been called for first time?
}

// NewBinaryStdIn creates the singleton BinaryStdIn struct
func NewBinaryStdIn() *BinaryStdIn {
	once.Do(func() {
		binaryStdIn = &BinaryStdIn{in: os.Stdin, buffer: 0, n: 0, isInitialized: false}
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
			bs.buffer = -1 // -1 means EOF
			bs.n = -1
		} else {
			panic(err)
		}
	} else {
		bs.n = 8
		bs.buffer = int(p[0])
	}
}

// IsEmpty returns true if standard input is empty.
func (bs *BinaryStdIn) IsEmpty() bool {
	if !bs.isInitialized {
		bs.initialize()
	}
	return bs.buffer == -1 // -1 means EOF
}

// ReadBool reads the next bit of data from standard input and return as a bool.
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

// ReadByte reads the next 8 bits from standard input and return as an 8-bit byte.
func (bs *BinaryStdIn) ReadByte() (byte, error) {
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")
	}
	// special case when aligned byte
	if bs.n == 8 {
		x := byte(bs.buffer)
		bs.fillBuffer()
		return x, nil
	}
	// combine last n bits of current buffer with first 8-n bits of new buffer
	x := bs.buffer
	x <<= (8 - bs.n) // filled with (8-bs.n) zero bits in the right
	oldN := bs.n
	bs.fillBuffer()
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")

	}
	bs.n = oldN
	// |= compound bitwise or operator  used with a variable and a constant
	// to "set" (set to 1) particular bits in a variable.
	//	x  x  x  x  x  x  x  x    variable
	//	0  0  0  0  0  0  1  1    mask
	//	----------------------
	//	x  x  x  x  x  x  1  1
	//  bits unchanged  |bits set
	x |= (bs.buffer >> bs.n)
	return byte(x), nil
}

// ReadInt reads the next 32 bits from standard input and return as a 32-bit int.
func (bs *BinaryStdIn) ReadInt() (int, error) {
	x := 0
	// 32 bit int equals 4 byte
	for i := 0; i < 4; i++ {
		b, err := bs.ReadByte()
		if err != nil {
			return -1, err
		}
		x <<= 8     // filled 8 zero bits in the right
		x |= int(b) // set the rightmost 8 bits to b
	}
	return x, nil
}

// ReadInt16 reads the next 16 bits from standard input and return as a 16-bit int16
func (bs *BinaryStdIn) ReadInt16() (int16, error) {
	var x int16 = 0
	for i := 0; i < 2; i++ {
		byteValue, err := bs.ReadByte()
		if err != nil {
			return -1, err
		}
		x <<= 8
		x |= int16(byteValue)
	}
	return x, nil
}

// ReadInt64 reads the next 64 bits from standard input and return as a 64-bit int64.
func (bs *BinaryStdIn) ReadInt64() (int64, error) {
	var x int64 = 0
	for i := 0; i < 2; i++ {
		intValue, err := bs.ReadInt()
		if err != nil {
			return -1, err
		}
		x <<= 32
		x |= int64(intValue)
	}
	return x, nil
}

// ReadIntR reads the next r bits from standard input and return as an r-bit int.
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

// ReadString reads the remaining bytes of data from standard input and return as a string.
func (bs *BinaryStdIn) ReadString() (string, error) {
	if bs.IsEmpty() {
		return "", errors.New("reading from empty input stream")
	}
	var s strings.Builder
	for !bs.IsEmpty() {
		b, _ := bs.ReadByte()
		s.WriteByte(b)
	}
	return s.String(), nil
}
