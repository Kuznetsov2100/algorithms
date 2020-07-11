package main

//   Reads in a binary file and writes out the bytes in hex, 16 per line.

//   % more abra.txt
//   ABRACADABRA!

//   % go run hexDump.go 16 < abra.txt
//   41 42 52 41 43 41 44 41 42 52 41 21
//   96 bits

//   Remark
//   --------------------------
//    - Similar to the Unix utilities od (octal dump) or hexdump (hexadecimal dump).

//   % od -t x1 < abra.txt
//   0000000 41 42 52 41 43 41 44 41 42 52 41 21
//   0000014

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/io/binarystdin"
)

func main() {
	bytesPerLine := 16
	var i int
	binarystdin := binarystdin.NewBinaryStdIn(os.Stdin)
	if len(os.Args) == 1 {
		bytesPerLine, _ = strconv.Atoi(os.Args[1])
	}

	for i = 0; !binarystdin.IsEmpty(); i++ {
		if bytesPerLine == 0 {
			//nolint:errcheck
			binarystdin.ReadByte()
			continue
		}
		if i == 0 {
			fmt.Printf("")
		} else if i%bytesPerLine == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Print(" ")
		}
		b, _ := binarystdin.ReadByte()
		fmt.Printf("%02x", b)
	}
	if bytesPerLine != 0 {
		fmt.Println()
	}
	fmt.Println(i*8, " bits")
}
