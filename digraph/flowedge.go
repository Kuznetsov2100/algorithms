package digraph

import (
	"fmt"
	"math"
)

// FlowEdge struct represents a capacitated edge with a flow in a FlowNetwork.
// Each edge consists of two integers (naming the two vertices), a real-valued capacity,
// and a real-valued flow. The data type provides methods for accessing the two endpoints of
// the directed edge and the weight. It also provides methods for changing the amount of flow
// on the edge and determining the residual capacity of the edge.
type FlowEdge struct {
	v        int     // from
	w        int     // to
	capacity float64 // capacity
	flow     float64 // flow
}

// NewFlowEdge initializes an edge from vertex v to vertex w with the given capacity and zero flow.
func NewFlowEdge(v, w int, capacity float64) *FlowEdge {
	if v < 0 {
		panic("vertex index must be a non-negative integer")
	}
	if w < 0 {
		panic("vertex index must be a non-negative integer")
	}
	if !(capacity >= 0.0) {
		panic("edge capacity must be non-negative")
	}
	return &FlowEdge{v: v, w: w, capacity: capacity, flow: 0.0}
}

// NewFlowEdgeF initializes an edge from vertex v to vertex w with the given capacity and flow.
func NewFlowEdgeF(v, w int, capacity, flow float64) *FlowEdge {
	if v < 0 {
		panic("vertex index must be a non-negative integer")
	}
	if w < 0 {
		panic("vertex index must be a non-negative integer")
	}
	if !(capacity >= 0.0) {
		panic("edge capacity must be non-negative")
	}
	if !(flow <= capacity) {
		panic("flow exceeds capacity")
	}
	if !(flow >= 0.0) {
		panic("flow must be non-negative")
	}
	return &FlowEdge{v: v, w: w, capacity: capacity, flow: flow}
}

// From returns the tail vertex of the edge.
func (fe *FlowEdge) From() int {
	return fe.v
}

// To returns the head vertex of the edge.
func (fe *FlowEdge) To() int {
	return fe.w
}

// Capacity returns the capacity of the edge.
func (fe *FlowEdge) Capacity() float64 {
	return fe.capacity
}

// Flow returns the flow on the edge.
func (fe *FlowEdge) Flow() float64 {
	return fe.flow
}

// Other returns the endpoint of the edge that is different from the given vertex
// (unless the edge represents a self-loop in which case it returns the same vertex).
func (fe *FlowEdge) Other(vertex int) int {
	if vertex == fe.v {
		return fe.w
	}
	if vertex == fe.w {
		return fe.v
	}
	panic("invalid endpoint")
}

// ResidualCapacityTo returns the residual capacity of the edge in the direction to the given vertex.
func (fe *FlowEdge) ResidualCapacityTo(vertex int) float64 {
	if vertex == fe.v {
		return fe.flow // backward edge
	}
	if vertex == fe.w {
		return fe.capacity - fe.flow // forward edge
	}
	panic("invalid endpoint")
}

// AddResidualFlowTo increases the flow on the edge in the direction to the given vertex.
func (fe *FlowEdge) AddResidualFlowTo(vertex int, delta float64) {
	if !(delta >= 0.0) {
		panic("Delta must be non-negative")
	}
	if vertex == fe.v {
		fe.flow -= delta // backward edge
	} else if vertex == fe.w {
		fe.flow += delta // forward edge
	} else {
		panic("invalid endpoint")
	}
	// round flow to 0 or capacity if within 1e-10
	if math.Abs(fe.flow) <= 1e-10 {
		fe.flow = 0
	}
	if math.Abs(fe.flow-fe.capacity) <= 1e-10 {
		fe.flow = fe.capacity
	}
	if !(fe.flow >= 0.0) {
		panic("flow is negative")
	}
	if !(fe.flow <= fe.capacity) {
		panic("flow exceeds capacity")
	}
}

// String returns a string representation of the edge.
func (fe *FlowEdge) String() string {
	return fmt.Sprintf("%d->%d %f/%f", fe.v, fe.w, fe.flow, fe.capacity)
}
