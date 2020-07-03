package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieST(t *testing.T) {
	assert := assert.New(t)

	shellsST := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	st := NewTrieST()

	for index, key := range shellsST {
		st.Put(key, index)
	}

	assert.Equal(7, st.Size())
	assert.False(st.IsEmpty())

	//Keys
	assert.Equal([]string{"by", "sea", "sells", "she", "shells", "shore", "the"}, st.Keys())

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

	//Get
	assert.Equal(0, st.Get("she").(int))

	assert.Nil(st.Get(""))

	assert.Nil(st.Get("love"))

	//Put
	st.Put("she", nil)
	assert.Equal([]string{"by", "sea", "sells", "shells", "shore", "the"}, st.Keys())

	st1 := NewTrieST()
	shellsST1 := []string{"she", "sells", "sea", "shells", "by", "the", "shores", "shore"}
	for index, key := range shellsST1 {
		st1.Put(key, index)
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
