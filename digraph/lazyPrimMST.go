package digraph

import (
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

// LazyPrimMST struct represents a data type for computing a minimum spanning tree in an edge-weighted
// graph. The edge weights can be positive, zero, or negative and need not be distinct.
// If the graph is not connected, it computes a minimum spanning forest, which is the union of minimum
// spanning trees in each connected component. The Weight() method returns the weight of a minimum
// spanning tree and the edges() method returns its edges.
// This implementation uses a lazy version of Prim's algorithm with a binary heap of edges.
// The constructor takes O(E log E) time in the worst case, where V is the number of vertices and E
// is the number of edges. Each instance method takes Θ(1) time. It uses O(E) extra space in the worst
// case (not including the edge-weighted graph).
type LazyPrimMST struct {
	weight float64              // total weight of MST
	mst    *arrayqueue.Queue    // edges in the MST
	marked []bool               // marked[v] = true if v on tree
	pq     *priorityqueue.MinPQ // edges with one endpoint in tree
}

// NewLazyPrimMST compute a minimum spanning tree (or forest) of an edge-weighted graph.
func NewLazyPrimMST(G *EdgeWeightedGraph) *LazyPrimMST {
	lp := &LazyPrimMST{
		mst:    arrayqueue.New(),
		pq:     priorityqueue.NewMinPQ(),
		marked: make([]bool, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		if !lp.marked[v] {
			lp.prim(G, v)
		}
	}
	return lp
}

func (lp *LazyPrimMST) prim(G *EdgeWeightedGraph, s int) {
	lp.scan(G, s)
	for !lp.pq.IsEmpty() {
		e, _ := lp.pq.DelMin()
		edge := e.(*Edge)
		v := edge.Either()
		w := edge.Other(v)
		if lp.marked[v] && lp.marked[w] {
			continue // lazy, both v and w already scanned
		}
		lp.mst.Enqueue(edge) // add e to MST
		lp.weight += edge.Weight()
		if !lp.marked[v] { // v becomes part of tree
			lp.scan(G, v)
		}
		if !lp.marked[w] { // w becomes part of tree
			lp.scan(G, w)
		}
	}
}

// add all edges e incident to v into pq if the other endpoint has not yet been scanned
func (lp *LazyPrimMST) scan(G *EdgeWeightedGraph, v int) {
	lp.marked[v] = true
	for _, e := range G.Adj(v) {
		if !lp.marked[e.Other(v)] {
			lp.pq.Insert(e)
		}
	}
}

// Edges returns the edges in a minimum spanning tree (or forest).
func (lp *LazyPrimMST) Edges() (edges []*Edge) {
	for _, x := range lp.mst.Values() {
		edges = append(edges, x.(*Edge))
	}
	return edges
}

// Weight returns the sum of the edge weights in a minimum spanning tree (or forest).
func (lp *LazyPrimMST) Weight() float64 {
	return lp.weight
}
