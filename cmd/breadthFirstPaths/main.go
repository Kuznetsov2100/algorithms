package main

//   %  go run main.go tinyCG.txt 0
//   0 to 0 (0):  0
//   0 to 1 (1):  0-1
//   0 to 2 (1):  0-2
//   0 to 3 (2):  0-2-3
//   0 to 4 (2):  0-2-4
//   0 to 5 (1):  0-5
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
	bfp := graph.NewBreadthFirstPaths(G, s)

	for v := 0; v < G.V(); v++ {
		if bfp.HasPathTo(v) {
			fmt.Printf("%d to %d (%d):  ", s, v, bfp.DistTo(v))
			for _, x := range bfp.PathTo(v) {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d (-): not connected\n", s, v)
		}
	}
}
