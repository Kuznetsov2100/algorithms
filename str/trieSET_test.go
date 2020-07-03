package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieSET(t *testing.T) {
	assert := assert.New(t)

	shellsST := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	st := NewTrieSET()
	ok10, err10 := st.Contains("love")
	assert.Nil(err10)
	assert.False(ok10)

	for _, key := range shellsST {
		//nolint:errcheck
		st.Add(key)
	}

	assert.Equal(7, st.Size())
	assert.False(st.IsEmpty())

	//Keys
	assert.Equal([]string{"by", "sea", "sells", "she", "shells", "shore", "the"}, st.Iterator())

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

	//Add
	err4 := st.Add("")
	assert.EqualError(err4, "first argument to Add() is empty string")

	st1 := NewTrieSET()
	shellsST1 := []string{"she", "sells", "sea", "shells", "by", "the", "shores", "shore"}
	for _, key := range shellsST1 {
		//nolint:errcheck
		st1.Add(key)
	}

	//Delete
	err5 := st1.Delete("")
	assert.EqualError(err5, "argument to Delete() is empty string")

	err5 = st1.Delete("love")
	assert.Nil(err5)

	err5 = st1.Delete("shores")
	assert.Nil(err5)

	err5 = st1.Delete("shore")
	assert.Nil(err5)
}
