package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	var textarray1, textarray2 []string
	whitespace := regexp.MustCompile(`\s+`)
	in1 := stdin.NewInFileLine(os.Args[1])
	for !in1.IsEmpty() {
		textarray1 = append(textarray1, in1.ReadString())
	}
	s := whitespace.ReplaceAllString(strings.TrimSpace(strings.Join(textarray1, "\n")), " ")

	in2 := stdin.NewInFileLine(os.Args[2])
	for !in2.IsEmpty() {
		textarray2 = append(textarray2, in2.ReadString())
	}
	t := whitespace.ReplaceAllString(strings.TrimSpace(strings.Join(textarray2, "\n")), " ")

	fmt.Printf("'%s'\n", str.Lcs(s, t))

}
