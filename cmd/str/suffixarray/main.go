package main

import (
	"fmt"
	"regexp"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(stdin.ReadAll(), " ")
	suffix := str.NewSuffixArray(s)

	fmt.Println("  i ind lcp rnk select")
	fmt.Println("----------------------")

	for i := 0; i < len(s); i++ {
		index := suffix.Index(i)
		ith := "\"" + s[index:min(index+50, len(s))] + "\""
		rank := suffix.Rank(s[index:])
		if i == 0 {
			fmt.Printf("%3d %3d %3s %3d %s\n", i, index, "-", rank, ith)
		} else {
			lcp := suffix.Lcp(i)
			fmt.Printf("%3d %3d %3d %3d %s\n", i, index, lcp, rank, ith)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
