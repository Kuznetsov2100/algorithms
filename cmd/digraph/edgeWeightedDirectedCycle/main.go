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
	finder := digraph.NewEdgeWeightedDirectedCycle(G)

	if finder.HasCycle() {
		fmt.Print("Directed cycle: ")
		for _, e := range finder.GetCycle() {
			fmt.Print(e, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("No directed cycle")
	}
	fmt.Println()
}
