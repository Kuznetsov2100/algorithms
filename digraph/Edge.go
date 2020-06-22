package digraph

import (
	"fmt"
	"math"
)

type Edge struct {
	v      int
	w      int
	weight float64
}

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

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Either() int {
	return e.v
}

func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	panic("illegal endpoint")
}

func (e *Edge) CompareTo(that *Edge) int {
	if e.weight < that.weight {
		return -1
	} else if e.weight > that.weight {
		return 1
	} else {
		return 0
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}
