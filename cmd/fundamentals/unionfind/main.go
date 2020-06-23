package main

//   Weighted quick-union by rank with path compression by halving.

//    % go run main.go < tinyUF.txt
//    4 3
//    3 8
//    6 5
//    9 4
//    2 1
//    5 0
//    7 2
//    6 1
//    2 components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/handane123/algorithms/fundamentals/unionfind"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	var uf *unionfind.UF
	stdin := stdin.NewStdInLine()

	if !stdin.IsEmpty() {
		n := stdin.ReadInt()

		uf = unionfind.NewUF(n)
	}
	for !stdin.IsEmpty() {
		numbers := strings.Split(stdin.ReadString(), " ")
		p, _ := strconv.Atoi(numbers[0])
		q, _ := strconv.Atoi(numbers[1])
		if uf.Find(p) == uf.Find(q) {
			continue
		}
		uf.Union(p, q)
		fmt.Println(p, " ", q)
	}
	fmt.Println(uf.Count(), " components")
}
