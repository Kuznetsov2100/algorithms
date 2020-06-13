package main

//  Run depth first search on an undirected graph.
//  Runs in O(E + V) time.
//
//  % java DepthFirstSearch tinyG.txt 0
//  0 1 2 3 4 5 6
//  NOT connected
//
//  % java DepthFirstSearch tinyG.txt 9
//  9 10 11 12
//  NOT connected

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := graph.NewGraphIn(in)
	s, _ := strconv.Atoi(os.Args[2])
	search := graph.NewDepthFirstSearch(G, s)
	for v := 0; v < G.V(); v++ {
		if search.IsMarked(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
	if search.Count() != G.V() {
		fmt.Println("NOT connected")
	} else {
		fmt.Println("connected")
	}
}
