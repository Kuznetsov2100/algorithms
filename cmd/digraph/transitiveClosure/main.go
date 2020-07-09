package main

//   Compute transitive closure of a digraph and support
//    reachability queries.

//    Preprocessing time: O(V(E + V)) time.
//    Query time: O(1).
//    Space: O(V^2).

//    % go run main.go tinyDG.txt
//           0  1  2  3  4  5  6  7  8  9 10 11 12
//    --------------------------------------------
//      0:   T  T  T  T  T  T
//      1:      T
//      2:   T  T  T  T  T  T
//      3:   T  T  T  T  T  T
//      4:   T  T  T  T  T  T
//      5:   T  T  T  T  T  T
//      6:   T  T  T  T  T  T  T        T  T  T  T
//      7:   T  T  T  T  T  T  T  T  T  T  T  T  T
//      8:   T  T  T  T  T  T  T  T  T  T  T  T  T
//      9:   T  T  T  T  T  T           T  T  T  T
//     10:   T  T  T  T  T  T           T  T  T  T
//     11:   T  T  T  T  T  T           T  T  T  T
//     12:   T  T  T  T  T  T           T  T  T  T

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)

	tc := digraph.NewTransitiveClosure(G)

	fmt.Print("     ")
	for v := 0; v < G.V(); v++ {
		fmt.Printf("%3d", v)
	}
	fmt.Println()
	fmt.Println("--------------------------------------------")

	for v := 0; v < G.V(); v++ {
		fmt.Printf("%3d: ", v)
		for w := 0; w < G.V(); w++ {
			if tc.Reachable(v, w) {
				fmt.Print("  T")
			} else {
				fmt.Print("   ")
			}

		}
		fmt.Println()
	}
}
