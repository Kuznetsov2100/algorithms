package main

// w1 * w2 * w3 * … * wn > 1
// log(w1) + log(w2) + log(w3) + … + log(wn) > 0
// (-log(w1)) + (-log(w2)) + (-log(w3)) + … + (-log(wn)) < 0

//   Arbitrage detection.

//   % more rates.txt
//   5
//   USD 1      0.741  0.657  1.061  1.005
//   EUR 1.349  1      0.888  1.433  1.366
//   GBP 1.521  1.126  1      1.614  1.538
//   CHF 0.942  0.698  0.619  1      0.953
//   CAD 0.995  0.732  0.650  1.049  1

//   % go run main.go rates.txt
//   1000.00000 USD =  741.00000 EUR
//    741.00000 EUR = 1012.20600 CAD
//   1012.20600 CAD = 1007.14497 USD

import (
	"fmt"
	"math"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	V := in.ReadInt()
	name := make([]string, V)
	G := digraph.NewEdgeWeightedDigraphV(V)

	for v := 0; v < V; v++ {
		in.Scanner.Scan() // because in.ReadString() do not forward the scan, so we do it manually.
		name[v] = in.ReadString()
		for w := 0; w < V; w++ {
			rate := in.ReadFloat64()
			e := digraph.NewDirectedEdge(v, w, -math.Log(rate))
			G.AddEdge(e)
		}
	}

	spt := digraph.NewBellmanFordSP(G, 0)
	if spt.HasNegativeCycle() {
		stake := 1000.0
		for _, e := range spt.NegativeCycle() {
			fmt.Printf("%10.5f %s ", stake, name[e.From()])
			stake *= math.Exp(-e.Weight())
			fmt.Printf("= %10.5f %s\n", stake, name[e.To()])
		}
	} else {
		fmt.Println("No arbitrage apportunity")
	}

}
