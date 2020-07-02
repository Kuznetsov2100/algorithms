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
		//nolint:errcheck
		st.Put(key, index)
	}

	assert.Equal(7, st.Size())
	assert.False(st.IsEmpty())

	//Keys
	assert.Equal([]string{"by", "sea", "sells", "she", "shells", "shore", "the"}, st.Keys())

	//Contains
	ok, err1 := st.Contains("")
	assert.False(ok)
	assert.EqualError(err1, "argument to Contains() is empty string")

	ok, err1 = st.Contains("by")
	assert.True(ok)
	assert.Nil(err1)

	//LongestPrefixOf
	actual, err := st.LongestPrefixOf("shellsort")
	assert.Nil(err)
	assert.Equal("shells", actual)

	actual, err = st.LongestPrefixOf("quicksort")
	assert.Nil(err)
	assert.Empty(actual)

	actual, err = st.LongestPrefixOf("")
	assert.Empty(actual)
	assert.EqualError(err, "argument to LongestPrefixOf() is empty string")

	actual, err = st.LongestPrefixOf("she")
	assert.Nil(err)
	assert.Equal("she", actual)

	//KeysWithPrefix
	assert.Equal([]string{"shore"}, st.KeysWithPrefix("shor"))

	//KeysThatMatch
	assert.Equal([]string{"shells"}, st.KeysThatMatch(".he.l."))

	//Get
	val, err3 := st.Get("she")
	assert.Nil(err3)
	assert.Equal(0, val.(int))

	val, err3 = st.Get("")
	assert.Nil(val)
	assert.EqualError(err3, "argument to Get() is empty string")

	val, err3 = st.Get("love")
	assert.Nil(err3)
	assert.Nil(val)

	//Put
	err4 := st.Put("", 10)
	assert.EqualError(err4, "first argument to Put() is empty string")

	err4 = st.Put("she", nil)
	assert.Nil(err4)
	assert.Equal([]string{"by", "sea", "sells", "shells", "shore", "the"}, st.Keys())

	//Delete
	st1 := NewTrieST()
	err5 := st1.Delete("")
	assert.EqualError(err5, "argument to Delete() is empty string")

	err5 = st.Delete("love")
	assert.Nil(err5)

	keys := st.Keys()
	for _, k := range keys {
		err6 := st.Delete(k)
		assert.Nil(err6)
	}
	assert.True(st.IsEmpty())
}
