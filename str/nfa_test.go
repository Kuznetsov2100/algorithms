package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNFA(t *testing.T) {
	assert := assert.New(t)

	nfa := NewNFA("((A*B|AC)D)")
	assert.True(nfa.Recognizes("AAAABD"))
	assert.False(nfa.Recognizes("AAAAC"))

	nfa = NewNFA("((a|(bc)*d)*)")
	assert.True(nfa.Recognizes("abcbcd"))
	assert.True(nfa.Recognizes("abcbcbcdaaaabcbcdaaaddd"))
	assert.PanicsWithValue("text contains the metacharacter '*'\n", func() { nfa.Recognizes("abc*d") })
	assert.PanicsWithValue("invalid regular expression", func() { NewNFA("((A*B|ACD)") })

	nfa = NewNFA("(ab+d+)")
	assert.True(nfa.Recognizes("abbbbbbdd"))
}
