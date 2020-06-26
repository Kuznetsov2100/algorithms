package digraph

import (
	"fmt"
	"math"

	"github.com/pkg/errors"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// AcyclicLP struct represents a data type for solving the single-source longest paths problem
// in edge-weighted directed acyclic graphs (DAGs). The edge weights can be positive, negative, or zero.
// This implementation uses a topological-sort based algorithm.
// The constructor takes O(V + E) time in the worst case, where V is the number of vertices
// and E is the number of edges. Each instance method takes O(1) time. It uses O(V) extra space
// (not including the edge-weighted digraph).
type AcyclicLP struct {
	distTo []float64       // distTo[v] = distance  of longest s->v path
	edgeTo []*DirectedEdge // edgeTo[v] = last edge on longest s->v path
}

// NewAcyclicLP computes a longest paths tree from s to every other vertex in the directed acyclic graph G.
func NewAcyclicLP(G *EdgeWeightedDigraph, s int) (*AcyclicLP, error) {
	lp := &AcyclicLP{
		distTo: make([]float64, G.V()),
		edgeTo: make([]*DirectedEdge, G.V()),
	}
	lp.validateVertex(s)
	for v := 0; v < G.V(); v++ {
		lp.distTo[v] = math.Inf(-1)
	}
	lp.distTo[s] = 0.0

	// visit vertices in topological order
	topological := NewTopologicalEWD(G)
	if !topological.HasOrder() {
		return nil, errors.New("digraph is not acyclic")
	}

	for _, v := range topological.Order() {
		for _, e := range G.Adj(v) {
			lp.relax(e)
		}
	}
	return lp, nil
}

// relax edge e, but update if you find a longer path
func (lp *AcyclicLP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if lp.distTo[w] < lp.distTo[v]+e.Weight() {
		lp.distTo[w] = lp.distTo[v] + e.Weight()
		lp.edgeTo[w] = e
	}
}

// DistTo returns the length of a longest path from the source vertex s to vertex v.
func (lp *AcyclicLP) DistTo(v int) float64 {
	lp.validateVertex(v)
	return lp.distTo[v]
}

// HasPathTo returns true if there is a path from the source vertex s to vertex v.
func (lp *AcyclicLP) HasPathTo(v int) bool {
	lp.validateVertex(v)
	return lp.distTo[v] > math.Inf(-1)
}

// PathTo returns a longest path from the source vertex s to vertex v.
func (lp *AcyclicLP) PathTo(v int) (edges []*DirectedEdge) {
	lp.validateVertex(v)
	if !lp.HasPathTo(v) {
		return nil
	}
	path := arraystack.New()
	for e := lp.edgeTo[v]; e != nil; e = lp.edgeTo[e.From()] {
		path.Push(e)
	}
	for _, e := range path.Values() {
		edges = append(edges, e.(*DirectedEdge))
	}
	return edges
}

func (lp *AcyclicLP) validateVertex(v int) {
	V := len(lp.distTo)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
