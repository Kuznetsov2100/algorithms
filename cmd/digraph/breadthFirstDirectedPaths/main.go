package main

//   %  go run main.go tinyDG.txt 3
//    3 to 0 (2):  3->2->0
//    3 to 1 (3):  3->2->0->1
//    3 to 2 (1):  3->2
//    3 to 3 (0):  3
//    3 to 4 (2):  3->5->4
//    3 to 5 (1):  3->5
//    3 to 6 (-):  not connected
//    3 to 7 (-):  not connected
//    3 to 8 (-):  not connected
//    3 to 9 (-):  not connected
//    3 to 10 (-):  not connected
//    3 to 11 (-):  not connected
//    3 to 12 (-):  not connected

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
	bfp := digraph.NewBreadthFirstDirectedPaths(G, s)

	for v := 0; v < G.V(); v++ {
		if bfp.HasPathTo(v) {
			fmt.Printf("%d to %d (%d):  ", s, v, bfp.DistTo(v))
			for _, x := range bfp.PathTo(v) {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Printf("->%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d (-): not connected\n", s, v)
		}
	}
}
