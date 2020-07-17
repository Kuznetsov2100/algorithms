package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLcs(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(" defg ", Lcs("abc defg ghjkl", "abc defg cfr ghjkl qwerty"))
	assert.Equal("abc", Lcs("abc", "abcd"))
	assert.Equal("defg", Lcs("defg", "defg"))
	assert.Equal("defg", Lcs("defgx", "defg"))
}
