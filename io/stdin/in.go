package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// In wraps a scanner using ScanWords as split function
type In struct {
	Scanner *bufio.Scanner
}

// NewInFileWords initializes an input stream from a local filename
func NewInFileWords(filename string) *In {
	File, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("can not open file: %s", filename))
	}
	scanner := bufio.NewScanner(File)
	scanner.Split(bufio.ScanWords)
	return &In{Scanner: scanner}
}

// NewInFileLine initializes an input stream from a local filename
func NewInFileLine(filename string) *In {
	File, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("can not open file: %s", filename))
	}
	scanner := bufio.NewScanner(File)
	return &In{Scanner: scanner}
}

// IsEmpty reports if the In is empty
func (in *In) IsEmpty() bool {
	return !in.Scanner.Scan()
}

// ReadString reads the next token and returns the string.
func (in *In) ReadString() string {
	return in.Scanner.Text()
}

// ReadInt reads the next token from this input stream, parses it as a int, and returns the int.
func (in *In) ReadInt() int {
	in.Scanner.Scan()
	v, err := strconv.Atoi(in.Scanner.Text())
	if err != nil {
		panic(err)
	}
	return v
}

// ReadFloat64 reads the next token from this input stream, parses it as a float64, and returns the float64.
func (in *In) ReadFloat64() float64 {
	in.Scanner.Scan()
	v, err := strconv.ParseFloat(in.Scanner.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}
