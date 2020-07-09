package stdin

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

var stdin *StdIn
var once sync.Once

// StdIn wraps the scanner
type StdIn struct {
	scanner *bufio.Scanner
}

// NewStdIn initialize Stdin using singleton mode
func NewStdIn() *StdIn {
	once.Do(func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		stdin = &StdIn{scanner}
	})
	return stdin
}

// NewStdInLine initialize Stdin using singleton mode
func NewStdInLine() *StdIn {
	once.Do(func() {
		scanner := bufio.NewScanner(os.Stdin)
		stdin = &StdIn{scanner}
	})
	return stdin
}

// IsEmpty reports if the In is empty
func (s *StdIn) IsEmpty() bool {
	return !s.scanner.Scan()
}

// ReadString reads the next token or the next line depends on splitfunc, and returns the string.
func (s *StdIn) ReadString() string {
	return s.scanner.Text()
}

// ReadInt reads the next token from this input stream, parses it as a int, and returns the int.
func (s *StdIn) ReadInt() int {
	v, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		panic(err)
	}
	return v
}

// ReadAllStrings reads all remaining tokens from standard input and returns them as a slice of strings.
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // split by words
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
