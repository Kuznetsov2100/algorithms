package digraph

import (
	"fmt"
	"math"

	"github.com/pkg/errors"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// AcyclicSP struct represents a data type for solving the single-source shortest paths problem
// in edge-weighted directed acyclic graphs (DAGs). The edge weights can be positive, negative, or zero.
// This implementation uses a topological-sort based algorithm. The constructor takes O(V + E) time
// in the worst case, where V is the number of vertices and E is the number of edges.
// Each instance method takes O(1) time. It uses O(V) extra space (not including the edge-weighted digraph).
type AcyclicSP struct {
	distTo []double        // distTo[v] = distance  of shortest s->v path
	edgeTo []*DirectedEdge // edgeTo[v] = last edge on shortest s->v path
}

// NewAcyclicSP computes a shortest paths tree from s to every other vertex in the directed acyclic graph G.
func NewAcyclicSP(G *EdgeWeightedDigraph, s int) (*AcyclicSP, error) {
	sp := &AcyclicSP{
		distTo: make([]double, G.V()),
		edgeTo: make([]*DirectedEdge, G.V()),
	}
	sp.validateVertex(s)
	for v := 0; v < G.V(); v++ {
		sp.distTo[v] = math.MaxFloat64
	}
	sp.distTo[s] = 0.0

	// visit vertices in topological order
	topological := NewTopologicalEWD(G)
	if !topological.HasOrder() {
		return nil, errors.New("digraph is not acyclic")
	}

	for _, v := range topological.Order() {
		for _, e := range G.Adj(v) {
			sp.relax(e)
		}
	}
	return sp, nil
}

func (sp *AcyclicSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if sp.distTo[w] > sp.distTo[v]+double(e.Weight()) {
		sp.distTo[w] = sp.distTo[v] + double(e.Weight())
		sp.edgeTo[w] = e
	}
}

// DistTo returns the length of a shortest path from the source vertex s to vertex v.
func (sp *AcyclicSP) DistTo(v int) float64 {
	sp.validateVertex(v)
	return float64(sp.distTo[v])
}

// HasPathTo returns true if there is a path from the source vertex s to vertex v.
func (sp *AcyclicSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo returns a shortest path from the source vertex s to vertex v.
func (sp *AcyclicSP) PathTo(v int) (edges []*DirectedEdge) {
	sp.validateVertex(v)
	if !sp.HasPathTo(v) {
		return nil
	}
	path := arraystack.New()
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path.Push(e)
	}
	for _, e := range path.Values() {
		edges = append(edges, e.(*DirectedEdge))
	}
	return edges
}

func (sp *AcyclicSP) validateVertex(v int) {
	V := len(sp.distTo)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
