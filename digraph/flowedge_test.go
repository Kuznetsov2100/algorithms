package digraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowEdge(t *testing.T) {
	assert := assert.New(t)

	// NewFlowEdge
	edge := NewFlowEdge(1, 2, 1.24)
	assert.Equal(1, edge.From())
	assert.Equal(2, edge.To())
	assert.Equal(1.24, edge.Capacity())
	assert.Equal(0.0, edge.Flow())
	assert.Equal(2, edge.Other(1))
	assert.Equal(1, edge.Other(2))
	assert.Equal("1->2 0.000000/1.240000", edge.String())
	assert.PanicsWithValue("invalid endpoint", func() { edge.Other(3) })

	assert.PanicsWithValue("vertex index must be a non-negative integer", func() { NewFlowEdge(-1, 2, 1.24) })
	assert.PanicsWithValue("vertex index must be a non-negative integer", func() { NewFlowEdge(1, -2, 1.24) })
	assert.PanicsWithValue("edge capacity must be non-negative", func() { NewFlowEdge(1, 2, -1.24) })

	// ResidualCapacityTo
	assert.Equal(0.0, edge.ResidualCapacityTo(1))
	assert.Equal(1.24, edge.ResidualCapacityTo(2))
	assert.PanicsWithValue("invalid endpoint", func() { edge.ResidualCapacityTo(3) })

	edge1 := NewFlowEdgeF(1, 2, 1.24, 1.1)
	assert.Equal(1, edge1.From())
	assert.Equal(2, edge1.To())
	assert.Equal(1.24, edge1.Capacity())
	assert.Equal(1.1, edge1.Flow())

	assert.PanicsWithValue("vertex index must be a non-negative integer", func() { NewFlowEdgeF(-1, 2, 1.24, 1.1) })
	assert.PanicsWithValue("vertex index must be a non-negative integer", func() { NewFlowEdgeF(1, -2, 1.24, 1.1) })
	assert.PanicsWithValue("edge capacity must be non-negative", func() { NewFlowEdgeF(1, 2, -1.24, 1.1) })
	assert.PanicsWithValue("flow exceeds capacity", func() { NewFlowEdgeF(1, 2, 1.24, 1.25) })
	assert.PanicsWithValue("flow must be non-negative", func() { NewFlowEdgeF(1, 2, 1.24, -1.1) })

	assert.PanicsWithValue("Delta must be non-negative", func() { edge1.AddResidualFlowTo(4, -1.3) })
	edge1.AddResidualFlowTo(2, 0.13)
	assert.Equal(1.23, edge1.Flow())
	edge1.AddResidualFlowTo(1, 1.2299999999999999999999999999)
	assert.Equal(0.0, edge1.Flow())
	assert.PanicsWithValue("invalid endpoint", func() { edge1.AddResidualFlowTo(3, 0.1) })

	edge2 := NewFlowEdgeF(1, 2, 1.24, 1.1)
	assert.PanicsWithValue("flow is negative", func() { edge2.AddResidualFlowTo(1, 1.2) })

	edge3 := NewFlowEdgeF(1, 2, 1.24, 1.1)
	assert.PanicsWithValue("flow exceeds capacity", func() { edge3.AddResidualFlowTo(2, 0.4) })

	edge4 := NewFlowEdgeF(1, 2, 1.14, 1.0)
	edge4.AddResidualFlowTo(2, 0.1399999999999999999)
	assert.Equal(1.14, edge4.Flow())
}
