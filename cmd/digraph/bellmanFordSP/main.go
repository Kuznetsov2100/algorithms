package main

//   Bellman-Ford shortest path algorithm. Computes the shortest path tree in
//    edge-weighted digraph G from vertex s, or finds a negative cost cycle
//    reachable from s.

//    % go run main.go tinyEWDn.txt 0
//    0 to 0 ( 0.00)
//    0 to 1 ( 0.93)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25   4->5  0.35   5->1  0.32
//    0 to 2 ( 0.26)  0->2  0.26
//    0 to 3 ( 0.99)  0->2  0.26   2->7  0.34   7->3  0.39
//    0 to 4 ( 0.26)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25
//    0 to 5 ( 0.61)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52   6->4 -1.25   4->5  0.35
//    0 to 6 ( 1.51)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52
//    0 to 7 ( 0.60)  0->2  0.26   2->7  0.34

//   % go run main.go tinyEWDnc.txt 0
//   4->5  0.35
//   5->4 -0.66

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

	sp := digraph.NewBellmanFordSP(G, s)

	if sp.HasNegativeCycle() {
		for _, e := range sp.NegativeCycle() {
			fmt.Println(e)
		}
	} else {
		for t := 0; t < G.V(); t++ {
			if sp.HasPathTo(t) {
				fmt.Printf("%d to %d (%5.2f)  ", s, t, sp.DistTo(t))
				for _, e := range sp.PathTo(t) {
					fmt.Print(e, "  ")
				}
				fmt.Println()
			} else {
				fmt.Printf("%d to %d     no path\n", s, t)
			}
		}
	}
}
