package digraph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedEdge(t *testing.T) {
	assert := assert.New(t)

	// v < 0
	assert.Panics(func() { NewDirectedEdge(-1, 2, 0.13) })

	// weight is NaN
	assert.Panics(func() { NewDirectedEdge(1, 2, math.NaN()) })

	edge := NewDirectedEdge(1, 2, 3.12)
	assert.Equal(1, edge.From())
	assert.Equal(2, edge.To())
	assert.Equal(3.12, edge.Weight())
	assert.Equal("1->2  3.12", edge.String())
}
