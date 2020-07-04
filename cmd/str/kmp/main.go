package main

//   Reads in two strings, the pattern and the input text, and
//    searches for the pattern in the input text using the
//    KMP algorithm.

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
	pattern := []byte(pat)
	text := []byte(txt)

	kmp1 := str.NewKMP(pat)
	offset1 := kmp1.Search(txt)

	kmp2 := str.NewKMPR(pattern, 256)
	offset2 := kmp2.SearchByte(text)

	fmt.Println("text:   ", txt)

	fmt.Print("pattern: ")
	for i := 0; i < offset1; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pat)

	fmt.Print("pattern: ")
	for i := 0; i < offset2; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pat)
}
