package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieSET(t *testing.T) {
	assert := assert.New(t)

	shellsST := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	st := NewTrieSET()
	assert.False(st.Contains("love"))

	for _, key := range shellsST {
		//nolint:errcheck
		st.Add(key)
	}

	assert.Equal(7, st.Size())
	assert.False(st.IsEmpty())

	//Keys
	assert.Equal([]string{"by", "sea", "sells", "she", "shells", "shore", "the"}, st.Iterator())

	//Contains
	assert.False(st.Contains(""))

	assert.True(st.Contains("by"))

	//LongestPrefixOf
	assert.Equal("shells", st.LongestPrefixOf("shellsort"))

	assert.Empty(st.LongestPrefixOf("quicksort"))

	assert.Empty(st.LongestPrefixOf(""))

	assert.Equal("she", st.LongestPrefixOf("she"))

	//KeysWithPrefix
	assert.Equal([]string{"shore"}, st.KeysWithPrefix("shor"))

	//KeysThatMatch
	assert.Equal([]string{"shells"}, st.KeysThatMatch(".he.l."))

	st1 := NewTrieSET()
	shellsST1 := []string{"she", "sells", "sea", "shells", "by", "the", "shores", "shore"}
	for _, key := range shellsST1 {
		st1.Add(key)
	}

	//Delete
	st1.Delete("")
	assert.Equal(8, st1.Size())

	st1.Delete("love")
	assert.Equal(8, st1.Size())

	st1.Delete("shores")
	assert.Equal(7, st1.Size())

	st1.Delete("shore")
	assert.Equal(6, st1.Size())
}
