package main

//   Compute a minimum spanning forest using a lazy version of Prim's algorithm.

//   %  go run main.go tinyEWG.txt
//   0-7 0.16000
//   1-7 0.19000
//   0-2 0.26000
//   2-3 0.17000
//   5-7 0.28000
//   4-5 0.35000
//   6-2 0.40000
//   1.81000

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G := digraph.NewEdgeWeightedGraphIn(in)
	mst := digraph.NewLazyPrimMST(G)
	for _, e := range mst.Edges() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f\n", mst.Weight())
}
