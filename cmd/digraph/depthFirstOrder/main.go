package main

//   Compute preorder and postorder for a digraph
//    Runs in O(E + V) time.

//    % go run main.go tinyDAG.txt
//       v   pre  post
//    ---------------
//       0    0    8
//       1    3    2
//       2    9   10
//       3   10    9
//       4    2    0
//       5    1    1
//       6    4    7
//       7   11   11
//       8   12   12
//       9    5    6
//      10    8    5
//      11    6    4
//      12    7    3
//    Preorder:  0 5 4 1 6 9 11 12 10 2 3 7 8
//    Postorder: 4 5 1 12 11 10 9 6 0 3 2 7 8
//    Reverse postorder: 8 7 2 3 0 6 9 10 11 12 1 5 4

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)

	dfo := digraph.NewDepthFirstOrder(G)
	fmt.Println("  v   pre   post")
	fmt.Println("----------------")
	for v := 0; v < G.V(); v++ {
		fmt.Printf("%4d %4d %4d\n", v, dfo.Pre(v), dfo.Post(v))
	}

	fmt.Print("Preorder:  ")
	for _, v := range dfo.PreOrder() {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Postorder:  ")
	for _, v := range dfo.PostOrder() {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Reverse postorder:  ")
	for _, v := range dfo.ReversePost() {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
