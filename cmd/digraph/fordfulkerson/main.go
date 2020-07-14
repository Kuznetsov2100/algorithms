package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
)

func main() {
	V, _ := strconv.Atoi(os.Args[1])
	E, _ := strconv.Atoi(os.Args[2])
	s := 0
	t := V - 1
	G := digraph.NewFlowNetworkVE(V, E)
	fmt.Println(G)

	maxflow := digraph.NewFordFulkerson(G, s, t)
	fmt.Println("max flow from ", s, " to ", t)
	for v := 0; v < G.V(); v++ {
		for _, e := range G.Adj(v) {
			if v == e.From() && e.Flow() > 0 {
				fmt.Println("   ", e)
			}
		}
	}

	fmt.Print("Min cut: ")
	for v := 0; v < G.V(); v++ {
		if maxflow.InCut(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	fmt.Println("max flow value = ", maxflow.Value())
}
