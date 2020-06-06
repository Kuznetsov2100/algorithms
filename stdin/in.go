package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// In wraps a scanner using ScanWords as split function
type In struct {
	scanner *bufio.Scanner
}

// NewInFileName initializes an input stream from a local filename
func NewInFileName(path string) *In {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic("can not open file: " + path)
	}
	scanner := bufio.NewScanner(inFile) // default SplitFunc: Scanlines
	return &In{scanner}
}

// IsEmpty reports if the In is empty
func (in *In) IsEmpty() bool {
	return !in.scanner.Scan()
}

// ReadString reads the next token and returns the string.
func (in *In) ReadString() string {
	return in.scanner.Text()
}

// ReadInt reads the next token from this input stream, parses it as a int, and returns the int.
func (in *In) ReadInt() int {
	in.scanner.Scan()
	v, _ := strconv.Atoi(in.scanner.Text())
	return v
}

// ReadFloat32 reads the next token from this input stream, parses it as a float32, and returns the float32.
func (in *In) ReadFloat32() float32 {
	in.scanner.Scan()
	v, _ := strconv.ParseFloat(in.scanner.Text(), 32)
	return float32(v)
}
