package main

//   A data type that computes the suffix array of a string.

//    % java SuffixArray < abra.txt
//     i ind lcp rnk  select
//    ---------------------------
//     0  11   -   0  "!"
//     1  10   0   1  "A!"
//     2   7   1   2  "ABRA!"
//     3   0   4   3  "ABRACADABRA!"
//     4   3   1   4  "ACADABRA!"
//     5   5   1   5  "ADABRA!"
//     6   8   0   6  "BRA!"
//     7   1   3   7  "BRACADABRA!"
//     8   4   0   8  "CADABRA!"
//     9   6   0   9  "DABRA!"
//    10   9   0  10  "RA!"
//    11   2   2  11  "RACADABRA!"

import (
	"fmt"
	"regexp"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	whitespace := regexp.MustCompile(`\s+`)
	s := whitespace.ReplaceAllString(stdin.ReadAll(), " ")
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
