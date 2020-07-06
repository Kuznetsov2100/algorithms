package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/str"
)

func main() {
	regexp := "(" + os.Args[1] + ")"
	txt := os.Args[2]
	nfa := str.NewNFA(regexp)
	fmt.Println(nfa.Recognizes(txt))
}
