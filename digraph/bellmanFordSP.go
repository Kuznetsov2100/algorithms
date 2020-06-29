package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// BellmanFordSP struct represents a data type for solving the single-source shortest paths problem
// in edge-weighted digraphs with no negative cycles. The edge weights can be positive, negative, or zero.
// This class finds either a shortest path from the source vertex s to every other vertex or
// a negative cycle reachable from the source vertex.
// This implementation uses a queue-based implementation of the Bellman-Ford-Moore algorithm.
// The constructor takes O(E V) time in the worst case, where V is the number of vertices and E
// is the number of edges. In practice, it performs much better. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the edge-weighted digraph).
type BellmanFordSP struct {
	distTo  []float64         // distTo[v] = distance  of shortest s->v path
	edgeTo  []*DirectedEdge   // edgeTo[v] = last edge on shortest s->v path
	onQueue []bool            // onQueue[v] = is v currently on the queue?
	queue   *arrayqueue.Queue // queue of vertices to relax
	cost    int               // number of calls to relax()
	cycle   []*DirectedEdge   // negative cycle (or nil if no such cycle)
}

// NewBellmanFordSP computes a shortest paths tree from s to every other vertex in the edge-weighted digraph G.
func NewBellmanFordSP(G *EdgeWeightedDigraph, s int) *BellmanFordSP {
	sp := &BellmanFordSP{
		distTo:  make([]float64, G.V()),
		edgeTo:  make([]*DirectedEdge, G.V()),
		onQueue: make([]bool, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		sp.distTo[v] = math.Inf(1)
	}
	sp.distTo[s] = 0.0

	sp.queue = arrayqueue.New()
	sp.queue.Enqueue(s)
	sp.onQueue[s] = true
	for !sp.queue.IsEmpty() && !sp.HasNegativeCycle() {
		val, _ := sp.queue.Dequeue()
		v := val.(int)
		sp.onQueue[v] = false
		sp.relax(G, v)
	}
	return sp
}

// relax vertex v and put other endpoints on queue if changed
func (sp *BellmanFordSP) relax(G *EdgeWeightedDigraph, v int) {
	for _, e := range G.Adj(v) {
		w := e.To()
		if sp.distTo[w] > sp.distTo[v]+e.Weight() {
			sp.distTo[w] = sp.distTo[v] + e.Weight()
			sp.edgeTo[w] = e
			// the only edges that could lead to a change in distTo[] are those leaving a vertex
			// whose distTo[] value changed in the previous pass. To keep track of such vertices,
			// we use a FIFO queue.
			if !sp.onQueue[w] {
				sp.queue.Enqueue(w)
				sp.onQueue[w] = true
			}
		}
		if sp.cost++; sp.cost%G.V() == 0 {
			sp.findNegativeCycle()
			if sp.HasNegativeCycle() {
				return // found a negative cycle
			}
		}
	}
}

// HasNegativeCycle returns true if there is a negative cycle reachable from the source vertex s
func (sp *BellmanFordSP) HasNegativeCycle() bool {
	return sp.cycle != nil
}

// NegativeCycle returns a negative cycle reachable from the source vertex s, or nil if there is no such cycle.
func (sp *BellmanFordSP) NegativeCycle() (edges []*DirectedEdge) {
	if !sp.HasNegativeCycle() {
		return nil
	}
	edges = make([]*DirectedEdge, len(sp.cycle))
	copy(edges, sp.cycle)
	return edges
}

// DistTo returns the length of a shortest path from the source vertex s to vertex v.
func (sp *BellmanFordSP) DistTo(v int) float64 {
	sp.validateVertex(v)
	return sp.distTo[v]
}

// HasPathTo returns true if there is a path from the source s to vertex v
func (sp *BellmanFordSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.Inf(1)
}

// PathTo returns a shortest path from the source s to vertex v.
func (sp *BellmanFordSP) PathTo(v int) (edges []*DirectedEdge) {
	sp.validateVertex(v)
	if sp.HasNegativeCycle() {
		panic("Negative cost cycle exists")
	}
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

func (sp *BellmanFordSP) findNegativeCycle() {
	V := len(sp.edgeTo)
	spt := NewEdgeWeightedDigraphV(V)
	for v := 0; v < V; v++ {
		if sp.edgeTo[v] != nil {
			spt.AddEdge(sp.edgeTo[v])
		}
	}
	finder := NewEdgeWeightedDirectedCycle(spt)
	sp.cycle = finder.GetCycle()
}

func (sp *BellmanFordSP) validateVertex(v int) {
	V := len(sp.distTo)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
