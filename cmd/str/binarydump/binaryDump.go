package main

//   Reads in a binary file and writes out the bits, n per line.

//   % more abra.txt
//   ABRACADABRA!

//   % go run binaryDump.go 16 < abra.txt
//   0100000101000010
//   0101001001000001
//   0100001101000001
//   0100010001000001
//   0100001001010010
//   0100000100100001
//   96 bits

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/io/binarystdin"
)

func main() {
	bitsPerLine := 16
	count := 0
	binarystdin := binarystdin.NewBinaryStdIn()
	if len(os.Args) == 1 {
		bitsPerLine, _ = strconv.Atoi(os.Args[1])
	}

	for count = 0; !binarystdin.IsEmpty(); count++ {
		if bitsPerLine == 0 {
			//nolint:errcheck
			binarystdin.ReadBool()
			continue
		} else if count != 0 && count%bitsPerLine == 0 {
			fmt.Println()
		}
		if ok, _ := binarystdin.ReadBool(); ok {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}

	}
	if bitsPerLine != 0 {
		fmt.Println()
	}
	fmt.Println(count, " bits")

}
