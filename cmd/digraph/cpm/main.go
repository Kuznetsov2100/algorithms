package main

//   Critical path method.

//   % go run main.go  jobsPC.txt
//    job   start  finish
//   --------------------
//      0     0.0    41.0
//      1    41.0    92.0
//      2   123.0   173.0
//      3    91.0   127.0
//      4    70.0   108.0
//      5     0.0    45.0
//      6    70.0    91.0
//      7    41.0    73.0
//      8    91.0   123.0
//      9    41.0    70.0
//   Finish time:   173.0

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {

	in := stdin.NewInFileWords(os.Args[1])
	n := in.ReadInt()

	source := 2 * n
	sink := 2*n + 1

	G := digraph.NewEdgeWeightedDigraphV(2*n + 2)
	for i := 0; i < n; i++ {
		duration := in.ReadFloat64()
		G.AddEdge(digraph.NewDirectedEdge(source, i, 0.0))
		G.AddEdge(digraph.NewDirectedEdge(i+n, sink, 0.0))
		G.AddEdge(digraph.NewDirectedEdge(i, i+n, duration))

		m := in.ReadInt()
		for j := 0; j < m; j++ {
			precedent := in.ReadInt()
			G.AddEdge(digraph.NewDirectedEdge(n+i, precedent, 0.0))
		}
	}

	lp, err := digraph.NewAcyclicLP(G, source)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("  job   start  finish")
	fmt.Println("---------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%4d %7.1f %7.1f\n", i, lp.DistTo(i), lp.DistTo(i+n))
	}
	fmt.Printf("Finish time: %7.1f\n", lp.DistTo(sink))

}
