package main

//   This program takes an RE as a command-line argument and prints
//    the lines from standard input having some substring that
//    is in the language described by the RE.

//    % more tinyL.txt
//    AC
//    AD
//    AAA
//    ABD
//    ADD
//    BCD
//    ABCCBD
//    BABAAA
//    BABBAAA

//    %  go run main.go "(A*B|AC)D" < tinyL.txt
//    ABD
//    ABCCBD

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	regexp := "(.*" + os.Args[1] + ".*)"
	nfa := str.NewNFA(regexp)
	stdin := stdin.NewStdInLine()
	for !stdin.IsEmpty() {
		line := stdin.ReadString()
		if nfa.Recognizes(line) {
			fmt.Println(line)
		}
	}
}
