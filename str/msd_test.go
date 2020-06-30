package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsdSort(t *testing.T) {
	assert := assert.New(t)

	original := []string{
		"she", "sells", "seashells", "by", "the",
		"sea", "shore", "the", "shells", "she", "sells",
		"are", "surely", "seashells",
	}

	expected := []string{
		"are", "by", "sea", "seashells", "seashells",
		"sells", "sells", "she", "she", "shells", "shore",
		"surely", "the", "the",
	}

	MsdSort(original)
	assert.Equal(expected, original)

	original1 := []string{
		"bed", "bug", "dad", "yes", "zoo",
		"now", "for", "tip", "ilk", "dim",
		"tag", "jot", "sob", "nob", "sky",
		"hut", "men", "egg", "few", "jay",
		"owl", "joy", "rap", "gig", "wee",
		"was", "wad", "fee", "tap", "tar",
		"dug", "jam", "all", "bad", "yet",
	}
	expected1 := []string{
		"all", "bad", "bed", "bug", "dad",
		"dim", "dug", "egg", "fee", "few",
		"for", "gig", "hut", "ilk", "jam",
		"jay", "jot", "joy", "men", "nob",
		"now", "owl", "rap", "sky", "sob",
		"tag", "tap", "tar", "tip", "wad",
		"was", "wee", "yes", "yet", "zoo",
	}

	MsdSort(original1)
	assert.Equal(expected1, original1)

	b := "hello"
	assert.PanicsWithValue("string index outbound", func() { charAt(b, 6) })
	assert.Equal(-1, charAt(b, 5))
}
