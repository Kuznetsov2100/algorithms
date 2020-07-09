package main

//   Compute the strongly-connected components of a digraph using the
//    Kosaraju-Sharir algorithm.

//    Runs in O(E + V) time.

//    % go run main.go tinyDG.txt
//    5 strong components
//    1
//    0 2 3 4 5
//    9 10 11 12
//    6 8
//    7

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)
	scc := digraph.NewKosarajuSharirSCC(G)

	m := scc.Count()
	fmt.Println(m, " strong components")

	components := make([]*arrayqueue.Queue, m)
	for i := 0; i < m; i++ {
		components[i] = arrayqueue.New()
	}

	for v := 0; v < G.V(); v++ {
		components[scc.Id(v)].Enqueue(v)
	}

	for i := 0; i < m; i++ {
		for _, v := range components[i].Values() {
			fmt.Print(v.(int), " ")
		}
		fmt.Println()
	}

}
