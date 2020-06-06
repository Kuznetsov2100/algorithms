package stdin

import (
	"bufio"
	"os"
)

// StdIn wraps the scanner
type StdIn struct {
	scanner *bufio.Scanner
}

// NewStdIn initialize Stdin
func NewStdIn() *StdIn {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &StdIn{scanner}
}

// IsEmpty reports if the In is empty
func (s *StdIn) IsEmpty() bool {
	return !s.scanner.Scan()
}

// ReadString reads the next token and returns the string.
func (s *StdIn) ReadString() string {
	return s.scanner.Text()
}

// ReadAllStrings reads all remaining tokens from standard input and returns them as a slice of strings.
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin) // default splitFunc:ScanLines
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
