package main

//  % go run main.go tinyCG.txt 0
//  0 to 0:  0
//  0 to 1:  0-2-1
//  0 to 2:  0-2
//  0 to 3:  0-2-3
//  0 to 4:  0-2-3-4
//  0 to 5:  0-2-3-5

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := graph.NewGraphIn(in)
	s, _ := strconv.Atoi(os.Args[2])
	dfp := graph.NewDepthFirstPaths(G, s)

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
