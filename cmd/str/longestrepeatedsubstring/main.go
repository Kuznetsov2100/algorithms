package main

//    Reads a text string from stdin, replaces all consecutive blocks of
//    whitespace with a single space, and then computes the longest
//    repeated substring in that text using a suffix array.

//    % go run main.go < tinyTale.txt
//    'st of times it was the '

//    % go run main.go < mobydick.txt
//    ',- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th'

//    % go run main.go
//    aaaaaaaaa
//    'aaaaaaaa'

//    % go run main.go
//    abcdefg
//    ''

import (
	"fmt"
	"regexp"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	whitespace := regexp.MustCompile(`\s+`)
	text := whitespace.ReplaceAllString(stdin.ReadAll(), " ")
	fmt.Printf("'%s'\n", str.Lrs(text))
}
