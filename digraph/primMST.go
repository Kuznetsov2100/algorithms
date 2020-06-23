package digraph

import (
	"math"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
)

// PrimMST struct represents a data type for computing a minimum spanning tree in an edge-weighted graph.
// The edge weights can be positive, zero, or negative and need not be distinct.
// If the graph is not connected, it computes a minimum spanning forest,
// which is the union of minimum spanning trees in each connected component.
// The weight() method returns the weight of a minimum spanning tree and the edges() method returns its edges.
// This implementation uses Prim's algorithm with an indexed binary heap.
// The constructor takes O(E log V) time in the worst case, where V is the number of vertices
// and E is the number of edges. Each instance method takes O(1) time. It uses O(V) extra space
// (not including the edge-weighted graph).
type PrimMST struct {
	edgeTo []*Edge  // edgeTo[v] = shortest edge from tree vertex to non-tree vertex
	distTo []double // distTo[v] = weight of shortest such edge
	marked []bool   // marked[v] = true if v on tree, false otherwise
	pq     *priorityqueue.IndexMinPQ
}

// NewPrimMST compute a minimum spanning tree (or forest) of an edge-weighted graph.
func NewPrimMST(G *EdgeWeightedGraph) *PrimMST {
	mst := &PrimMST{
		edgeTo: make([]*Edge, G.V()),
		distTo: make([]double, G.V()),
		marked: make([]bool, G.V()),
		pq:     priorityqueue.NewIndexMinPQ(G.V()),
	}
	for i := range mst.distTo {
		mst.distTo[i] = math.MaxFloat64
	}

	for v := 0; v < G.V(); v++ {
		if !mst.marked[v] {
			mst.prim(G, v)
		}
	}
	return mst
}

// run Prim's algorithm in graph G, starting from vertex s
func (mst *PrimMST) prim(G *EdgeWeightedGraph, s int) {
	mst.distTo[s] = 0.0
	//nolint:errcheck
	mst.pq.Insert(s, mst.distTo[s])
	for !mst.pq.IsEmpty() {
		v, _ := mst.pq.DelMin() // Add closest vertex to tree.
		mst.scan(G, v)
	}
}

// scan vertex v
func (mst *PrimMST) scan(G *EdgeWeightedGraph, v int) {
	mst.marked[v] = true
	for _, e := range G.Adj(v) {
		w := e.Other(v)
		if mst.marked[w] {
			continue // v-w is obsolete edge
		}
		if e.Weight() < float64(mst.distTo[w]) {
			// Edge e is new best connection from tree to w.
			mst.distTo[w] = double(e.Weight())
			mst.edgeTo[w] = e
			if mst.pq.Contains(w) {
				//nolint:errcheck
				mst.pq.DecreaseKey(w, mst.distTo[w])
			} else {
				//nolint:errcheck
				mst.pq.Insert(w, mst.distTo[w])
			}
		}
	}
}

// Edges returns the edges in a minimum spanning tree (or forest).
func (mst *PrimMST) Edges() (edges []*Edge) {
	for v := 0; v < len(mst.edgeTo); v++ {
		if e := mst.edgeTo[v]; e != nil {
			edges = append(edges, e)
		}
	}
	return edges
}

// Weight returns the sum of the edge weights in a minimum spanning tree (or forest).
func (mst *PrimMST) Weight() float64 {
	weight := 0.0
	for _, e := range mst.Edges() {
		weight += e.Weight()
	}
	return weight
}

type double float64

func (d double) CompareTo(k priorityqueue.Key) int {
	that := k.(double)
	if d < that {
		return -1
	} else if d > that {
		return 1
	} else {
		return 0
	}
}
