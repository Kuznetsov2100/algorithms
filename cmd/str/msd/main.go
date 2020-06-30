package main

//   Sort an array of strings or integers using MSD radix sort.

//    % go run main.go < shells.txt
//    are
//    by
//    sea
//    seashells
//    seashells
//    sells
//    sells
//    she
//    she
//    shells
//    shore
//    surely
//    the
//    the

import (
	"fmt"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	a := stdin.ReadAllStrings()
	n := len(a)

	str.MsdSort(a)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}
