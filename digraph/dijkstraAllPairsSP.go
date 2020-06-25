package digraph

import (
	"fmt"
	"math"
)

// DijkstraAllPairsSP struct represents a data type for solving the all-pairs shortest paths problem
// in edge-weighted digraphs where the edge weights are nonnegative.
// This implementation runs Dijkstra's algorithm from each vertex. The constructor takes O(V (E log V)) time
// in the worst case, where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V2) extra space (not including the edge-weighted digraph).
type DijkstraAllPairsSP struct {
	all []*DijkstraSP
}

// NewDijkstraAllPairsSP computes a shortest paths tree from each vertex to to every other vertex in the edge-weighted digraph G.
func NewDijkstraAllPairsSP(G *EdgeWeightedDigraph) *DijkstraAllPairsSP {
	pairs := &DijkstraAllPairsSP{all: make([]*DijkstraSP, G.V())}
	for v := 0; v < G.V(); v++ {
		pairs.all[v] = NewDijkstraSP(G, v)
	}
	return pairs
}

// Path returns a shortest path from vertex s to vertex t.
func (pairs *DijkstraAllPairsSP) Path(s, t int) (edges []*DirectedEdge) {
	pairs.validateVertex(s)
	pairs.validateVertex(t)
	return pairs.all[s].PathTo(t)
}

// HasPath returns true if there is a path from the vertex s to vertex t
func (pairs *DijkstraAllPairsSP) HasPath(s, t int) bool {
	pairs.validateVertex(s)
	pairs.validateVertex(t)
	return pairs.all[s].DistTo(t) < math.MaxFloat64
}

// Dist returns the length of a shortest path from vertex s to vertex t.
func (pairs *DijkstraAllPairsSP) Dist(s, t int) float64 {
	pairs.validateVertex(s)
	pairs.validateVertex(t)
	return pairs.all[s].DistTo(t)
}

func (pairs *DijkstraAllPairsSP) validateVertex(v int) {
	V := len(pairs.all)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
