package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)
	s, _ := strconv.Atoi(os.Args[2])
	dfp := digraph.NewDepthFirstDirectedPaths(G, s)
	for v := 0; v < G.V(); v++ {
		if dfp.HasPathTo(v) {
			fmt.Printf("%d to %d:  ", s, v)
			for _, x := range dfp.PathTo(v) {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d: not connected\n", s, v)
		}
	}
}
