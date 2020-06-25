package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G := digraph.NewEdgeWeightedDigraphIn(in)
	s, _ := strconv.Atoi(os.Args[2])

	sp := digraph.NewDijkstraSP(G, s)

	for t := 0; t < G.V(); t++ {
		if sp.HasPathTo(t) {
			fmt.Printf("%d to %d (%.2f)  ", s, t, sp.DistTo(t))
			for _, e := range sp.PathTo(t) {
				fmt.Print(e, "  ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d     no path\n", s, t)
		}
	}

}
