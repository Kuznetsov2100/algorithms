package main

//   Computes longeset paths in an edge-weighted acyclic digraph.

//   Remark: should probably check that graph is a DAG before running

//   % go run main.go tinyEWDAG.txt 5
//   5 to 0 (2.44)  5->1  0.32   1->3  0.29   3->6  0.52   6->4  0.93   4->0  0.38
//   5 to 1 (0.32)  5->1  0.32
//   5 to 2 (2.77)  5->1  0.32   1->3  0.29   3->6  0.52   6->4  0.93   4->7  0.37   7->2  0.34
//   5 to 3 (0.61)  5->1  0.32   1->3  0.29
//   5 to 4 (2.06)  5->1  0.32   1->3  0.29   3->6  0.52   6->4  0.93
//   5 to 5 (0.00)
//   5 to 6 (1.13)  5->1  0.32   1->3  0.29   3->6  0.52
//   5 to 7 (2.43)  5->1  0.32   1->3  0.29   3->6  0.52   6->4  0.93   4->7  0.37

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

	lp, err := digraph.NewAcyclicLP(G, s)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	for t := 0; t < G.V(); t++ {
		if lp.HasPathTo(t) {
			fmt.Printf("%d to %d (%.2f)  ", s, t, lp.DistTo(t))
			for _, e := range lp.PathTo(t) {
				fmt.Print(e, "  ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d     no path\n", s, t)
		}
	}
}
