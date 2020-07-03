package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTST(t *testing.T) {
	assert := assert.New(t)

	shellsST := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	st := NewTST()
	//KeysWithPrefix
	assert.Equal([]string(nil), st.KeysWithPrefix("se"))

	for index, key := range shellsST {
		st.Put(key, index)
	}

	assert.Equal(7, st.Size())
	assert.False(st.IsEmpty())

	//Keys
	assert.Equal([]string{"by", "sea", "sells", "she", "shells", "shore", "the"}, st.Keys())

	//Contains
	assert.True(st.Contains("by"))

	//LongestPrefixOf
	assert.Equal("shells", st.LongestPrefixOf("shellsort"))

	assert.Empty(st.LongestPrefixOf("quicksort"))

	assert.Empty(st.LongestPrefixOf(""))

	assert.Equal("she", st.LongestPrefixOf("she"))

	//KeysWithPrefix
	assert.Equal([]string{"shore"}, st.KeysWithPrefix("shor"))

	assert.Equal([]string{"sea"}, st.KeysWithPrefix("sea"))

	//KeysThatMatch
	assert.Equal([]string{"shells"}, st.KeysThatMatch(".he.l."))

	//Get
	assert.Equal(0, st.Get("she").(int))

	assert.Nil(st.Get("love"))

	assert.PanicsWithValue("key must have length >= 1", func() { st.Get("") })

	//Put
	st.Put("sea", nil)
	assert.Equal(6, st.Size())
}
