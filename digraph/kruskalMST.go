package digraph

import (
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/fundamentals/unionfind"
)

// KruskalMST struct represents a data type for computing a minimum spanning tree in an edge-weighted graph.
// This implementation uses Krusal's algorithm and the union-find data type.
// The constructor takes O(E log E) time in the worst case. Each instance method takes O(1) time.
// It uses O(E) extra space (not including the graph).
type KruskalMST struct {
	weight float64 // weight of MST
	mst    []*Edge // edges in MST
}

// NewKruskalMST compute a minimum spanning tree (or forest) of an edge-weighted graph.
func NewKruskalMST(G *EdgeWeightedGraph) *KruskalMST {
	kr := &KruskalMST{}
	// more efficient to build heap by passing array of edges
	pq := priorityqueue.NewMinPQ()
	for _, e := range G.Edges() {
		pq.Insert(e)
	}

	uf := unionfind.NewUF(G.V())
	for !pq.IsEmpty() && len(kr.mst) < (G.V()-1) {
		e, _ := pq.DelMin()
		edge := e.(*Edge)
		v := edge.Either()
		w := edge.Other(v)
		if uf.Find(v) != uf.Find(w) { // v-w does not create a cycless
			uf.Union(v, w)                // merge v and w components
			kr.mst = append(kr.mst, edge) // add edge e to mst
			kr.weight += edge.Weight()
		}
	}

	return kr
}

// Edges returns the edges in a minimum spanning tree (or forest).
func (kr *KruskalMST) Edges() (edges []*Edge) {
	edges = make([]*Edge, len(kr.mst))
	copy(edges, kr.mst)
	return edges
}

// Weight returns the sum of the edge weights in a minimum spanning tree (or forest).
func (kr *KruskalMST) Weight() float64 {
	return kr.weight
}
