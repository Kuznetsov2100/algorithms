package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/handane123/algorithms/graph"
)

func main() {
	V1, _ := strconv.Atoi(os.Args[1])
	V2, _ := strconv.Atoi(os.Args[2])
	E, _ := strconv.Atoi(os.Args[3])
	F, _ := strconv.Atoi(os.Args[4])

	G, _ := graph.BipartiteGraph(V1, V2, E)
	for i := 0; i < F; i++ {
		v := rand.Intn(V1 + V2)
		w := rand.Intn(V1 + V2)
		G.AddEdge(v, w)
	}
	fmt.Println(G)

	b := graph.NewBipartite(G)

	if b.IsBipartite() {
		fmt.Println("Graph is bipartite")
		for v := 0; v < G.V(); v++ {
			fmt.Println(v, ": ", b.Color(v))
		}
	} else {
		fmt.Println("Graph has an odd-length cycle: ")
		for _, x := range b.OddCycle() {
			fmt.Print(x, " ")
		}
		fmt.Println()
	}
}
