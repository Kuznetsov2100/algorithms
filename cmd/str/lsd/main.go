package main

//   LSD radix sort

//  - Sort a []String array of n extended ASCII strings (R = 256), each of length w.

//    Uses extra space proportional to n + R.

//    % go run main.go < words3.txt
//    all
//    bad
//    bed
//    bug
//    dad
//    ...
//    yes
//    yet
//    zoo

import (
	"fmt"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	a := stdin.ReadAllStrings()
	n := len(a)
	w := len(a[0])

	str.LsdSort(a, w)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}
