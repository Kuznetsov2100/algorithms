package main

//   Computes shortest paths in an edge-weighted acyclic digraph.

//    % go run main.go tinyEWDAG.txt 5
//    5 to 0 (0.73)  5->4  0.35   4->0  0.38
//    5 to 1 (0.32)  5->1  0.32
//    5 to 2 (0.62)  5->7  0.28   7->2  0.34
//    5 to 3 (0.61)  5->1  0.32   1->3  0.29
//    5 to 4 (0.35)  5->4  0.35
//    5 to 5 (0.00)
//    5 to 6 (1.13)  5->1  0.32   1->3  0.29   3->6  0.52
//    5 to 7 (0.28)  5->7  0.28

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G := digraph.NewEdgeWeightedDigraphIn(in)
	s, _ := strconv.Atoi(os.Args[2])

	sp, err := digraph.NewAcyclicSP(G, s)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

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
