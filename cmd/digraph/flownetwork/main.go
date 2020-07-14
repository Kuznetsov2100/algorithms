package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {

	in := stdin.NewInFileWords(os.Args[1])
	G := digraph.NewFlowNetworkIn(in)
	fmt.Println(G)
}
