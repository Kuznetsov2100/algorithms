package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/str"
)

func main() {
	huffman := str.NewHuffman(os.Stdin, os.Stdout)
	if os.Args[1] == "-" {
		huffman.Compress()
	} else if os.Args[1] == "+" {
		huffman.Expand()
	} else {
		fmt.Println("illegal command line argument")
	}
}
