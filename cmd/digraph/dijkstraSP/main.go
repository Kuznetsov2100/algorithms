package main

//   Dijkstra's algorithm. Computes the shortest path tree.
//    Assumes all weights are nonnegative.

//    % go run main.go tinyEWD.txt 0
//    0 to 0 (0.00)
//    0 to 1 (1.05)  0->4  0.38   4->5  0.35   5->1  0.32
//    0 to 2 (0.26)  0->2  0.26
//    0 to 3 (0.99)  0->2  0.26   2->7  0.34   7->3  0.39
//    0 to 4 (0.38)  0->4  0.38
//    0 to 5 (0.73)  0->4  0.38   4->5  0.35
//    0 to 6 (1.51)  0->2  0.26   2->7  0.34   7->3  0.39   3->6  0.52
//    0 to 7 (0.60)  0->2  0.26   2->7  0.34

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
