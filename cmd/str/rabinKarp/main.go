package main

//    Reads in two strings, the pattern and the input text, and
//    searches for the pattern in the input text using the
//    bad-character rule part of the Boyer-Moore algorithm.
//    (does not implement the strong good suffix rule)

//    % go run main.go abracadabra abacadabrabracabracadabrabrabracad
//    text:    abacadabrabracabracadabrabrabracad
//    pattern:               abracadabra

//    % go run main.go rab abacadabrabracabracadabrabrabracad
//    text:    abacadabrabracabracadabrabrabracad
//    pattern:         rab

//    % go run main.go bcara abacadabrabracabracadabrabrabracad
//    text:    abacadabrabracabracadabrabrabracad
//    pattern:                                   bcara

//    % go run main.go rabrabracad abacadabrabracabracadabrabrabracad
//    text:    abacadabrabracabracadabrabrabracad
//    pattern:                        rabrabracad

//    % go run main.go abacad abacadabrabracabracadabrabrabracad
//    text:    abacadabrabracabracadabrabrabracad
//    pattern: abacad

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/str"
)

func main() {
	pat := os.Args[1]
	txt := os.Args[2]

	rk := str.NewRabinKarp(pat)
	offset := rk.Search(txt)

	fmt.Println("text:   ", txt)

	fmt.Print("pattern: ")
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pat)
}
