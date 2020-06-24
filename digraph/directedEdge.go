package digraph

import (
	"fmt"
	"math"
)

// DirectedEdge struct represents a weighted edge in an EdgeWeightedDigraph.
// Each edge consists of two integers (naming the two vertices) and a real-value weight.
// The data type provides methods for accessing the two endpoints of the directed edge and the weight.
type DirectedEdge struct {
	v      int
	w      int
	weight float64
}

// NewDirectedEdge initializes a directed edge from vertex v to vertex w with the given weight.
func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	if v < 0 || w < 0 {
		panic("vertex names must be non negative")
	}
	if math.IsNaN(weight) {
		panic("weight is NaN")
	}
	return &DirectedEdge{v: v, w: w, weight: weight}
}

// From returns the tail vertex of the directed edge.
func (de *DirectedEdge) From() int {
	return de.v
}

// To returns the head vertex of the directed edge.
func (de *DirectedEdge) To() int {
	return de.w
}

// Weight returns the weight of the directed edge.
func (de *DirectedEdge) Weight() float64 {
	return de.weight
}

// String returns a string representation of the directed edge.
func (de *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %5.2f", de.v, de.w, de.weight)
}
