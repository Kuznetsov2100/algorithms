package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuick3string(t *testing.T) {
	assert := assert.New(t)

	original := []string{
		"bed", "bug", "dad", "yes", "zoo",
		"now", "for", "tip", "ilk", "dim",
		"tag", "jot", "sob", "nob", "sky",
		"hut", "men", "egg", "few", "jay",
		"owl", "joy", "rap", "gig", "wee",
		"was", "wad", "fee", "tap", "tar",
		"dug", "jam", "all", "bad", "yet",
	}
	expected := []string{
		"all", "bad", "bed", "bug", "dad",
		"dim", "dug", "egg", "fee", "few",
		"for", "gig", "hut", "ilk", "jam",
		"jay", "jot", "joy", "men", "nob",
		"now", "owl", "rap", "sky", "sob",
		"tag", "tap", "tar", "tip", "wad",
		"was", "wee", "yes", "yet", "zoo",
	}

	original1 := []string{
		"bed", "bug", "dad", "yes", "zoo",
		"now", "for", "tip", "ilk", "dim",
		"tag", "jot", "sob", "nob", "sky",
		"hut", "men", "egg", "jam", "jay",
		"owl", "joy", "rap", "gig", "wee",
		"was", "wad", "fee", "tap", "tar",
		"dug", "few", "all", "bad", "yet",
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

	Quick3string(original)
	assert.Equal(expected, original)

	Quick3string(original1)
	assert.Equal(expected1, original1)
}
