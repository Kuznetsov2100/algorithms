package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
)

// Edge struct represents a weighted edge in an EdgeWeightedGraph.
// Each edge consists of two integers (naming the two vertices) and a real-value weight.
// The data type provides methods for accessing the two endpoints of the edge and the weight.
// The natural order for this data type is by ascending order of weight.
type Edge struct {
	v      int
	w      int
	weight float64
}

/// NewEdge initializes an edge between vertices v and w of the given weight.
func NewEdge(v, w int, weight float64) *Edge {
	if v < 0 {
		panic("vertex index must be non negative integer")
	}
	if w < 0 {
		panic("vertex index must be non negative integer")
	}
	if math.IsNaN(weight) {
		panic("weight is NaN")
	}
	return &Edge{v: v, w: w, weight: weight}
}

// Weight returns the weight of this edge.
func (e *Edge) Weight() float64 {
	return e.weight
}

// Either returns either endpoint of this edge.
func (e *Edge) Either() int {
	return e.v
}

// Other returns the endpoint of this edge that is different from the given vertex.
func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	panic("illegal endpoint")
}

// CompareTo compares two edges by weight.
func (e *Edge) CompareTo(k priorityqueue.Key) int {
	that := k.(*Edge)
	if e.weight < that.weight {
		return -1
	} else if e.weight > that.weight {
		return 1
	} else {
		return 0
	}
}

// String returns a string representation of this edge.
func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}
