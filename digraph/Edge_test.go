package digraph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	assert := assert.New(t)

	// v < 0
	assert.Panics(func() { NewEdge(-1, 2, 3.64) })

	// w < 0
	assert.Panics(func() { NewEdge(3, -2, 5.32) })

	// weight is NaN
	assert.Panics(func() { NewEdge(4, 5, math.NaN()) })

	edge := NewEdge(9, 8, 1.24)
	// String
	assert.Equal("9-8 1.24000", edge.String())
	// Weight
	assert.Equal(1.24, edge.Weight())
	// Either
	assert.Equal(9, edge.Either())

	// Other
	assert.Equal(8, edge.Other(9))
	assert.Equal(9, edge.Other(8))
	assert.Panics(func() { edge.Other(3) })

	// CompareTo
	that := NewEdge(1, 2, 4.56)
	assert.Equal(-1, edge.CompareTo(that))
	that = NewEdge(1, 2, 0.82)
	assert.Equal(1, edge.CompareTo(that))
	that = NewEdge(9, 8, 1.24)
	assert.Equal(0, edge.CompareTo(that))

}
