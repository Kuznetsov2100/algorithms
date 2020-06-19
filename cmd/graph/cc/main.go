package main

//   % go main.go tinyG.txt
//   3 components
//   0 1 2 3 4 5 6
//   7 8
//   9 10 11 12

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := graph.NewGraphIn(in)
	cc := graph.NewCC(G)

	m := cc.Count()
	fmt.Println(m, " components")

	components := make([]*arrayqueue.Queue, m)
	for i := 0; i < m; i++ {
		components[i] = arrayqueue.New()
	}

	for v := 0; v < G.V(); v++ {
		components[cc.Id(v)].Enqueue(v)
	}

	for i := 0; i < m; i++ {
		for _, v := range components[i].Values() {
			fmt.Print(v.(int), " ")
		}
		fmt.Println()
	}
}
