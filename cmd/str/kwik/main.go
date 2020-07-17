package main

//   Keyword-in-context search.

//   % go run main.go tale.txt 15
//   majesty
//    most gracious majesty king george th
//   rnkeys and the majesty of the law fir
//   on against the majesty of the people
//   se them to his majestys chief secreta
//   h lists of his majestys forces and of

//   the worst
//   w the best and the worst are known to y
//   f them give me the worst first there th
//   for in case of the worst is a friend in
//   e roomdoor and the worst is over then a
//   pect mr darnay the worst its the wisest
//   is his brother the worst of a bad race
//   ss in them for the worst of health for
//    you have seen the worst of her agitati
//   cumwented into the worst of luck buuust
//   n your brother the worst of the bad rac
//    full share in the worst of the day pla
//   mes to himself the worst of the strife
//   f times it was the worst of times it wa
//   ould hope that the worst was over well
//   urage business the worst will be over i
//   clesiastics of the worst world worldly

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	var textarray []string
	whitespace := regexp.MustCompile(`\s+`)
	in := stdin.NewInFileLine(os.Args[1])
	for !in.IsEmpty() {
		textarray = append(textarray, in.ReadString())
	}
	text := whitespace.ReplaceAllString(strings.Join(textarray, "\n"), " ")
	n := len(text)

	context, _ := strconv.Atoi(os.Args[2])

	sa := str.NewSuffixArray(text)
	stdin := stdin.NewStdInLine()
	for !stdin.IsEmpty() {
		query := strings.TrimSpace(stdin.ReadString())
		for i := sa.Rank(query); i < n; i++ {
			from1 := sa.Index(i)
			to1 := min(n, from1+len(query))
			if query != text[from1:to1] {
				break
			}
			from2 := max(0, sa.Index(i)-context)
			to2 := min(n, sa.Index(i)+context+len(query))
			fmt.Println(text[from2:to2])
		}
		fmt.Println()
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
