package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/str"
)

func main() {
	lzw := str.NewLZW(os.Stdin, os.Stdout)
	if os.Args[1] == "-" {
		lzw.Compress()
	} else if os.Args[1] == "+" {
		lzw.Expand()
	} else {
		fmt.Println("illegal command line argument")
	}
}
