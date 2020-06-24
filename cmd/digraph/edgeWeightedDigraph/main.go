package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G := digraph.NewEdgeWeightedDigraphIn(in)
	fmt.Println(G)
}
