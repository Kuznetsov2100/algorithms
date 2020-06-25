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

	spt := digraph.NewDijkstraAllPairsSP(G)

	fmt.Printf("  ")
	for v := 0; v < G.V(); v++ {
		fmt.Printf("%6d ", v)
	}
	fmt.Println()
	for v := 0; v < G.V(); v++ {
		fmt.Printf("%3d: ", v)
		for w := 0; w < G.V(); w++ {
			if spt.HasPath(v, w) {
				fmt.Printf("%6.2f ", spt.Dist(v, w))
			} else {
				fmt.Printf("  Inf ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	for v := 0; v < G.V(); v++ {
		for w := 0; w < G.V(); w++ {
			if spt.HasPath(v, w) {
				fmt.Printf("%d to %d (%5.2f)  ", v, w, spt.Dist(v, w))
				for _, e := range spt.Path(v, w) {
					fmt.Print(e, "  ")
				}
				fmt.Println()
			} else {
				fmt.Printf("%d to %d no path\n", v, w)
			}
		}
	}
}
