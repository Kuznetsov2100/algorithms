package str

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLrs(t *testing.T) {
	assert := assert.New(t)

	tinyTale := "it was the best of times it was the worst of times\n" +
		"it was the age of wisdom it was the age of foolishness\n" +
		"it was the epoch of belief it was the epoch of incredulity\n" +
		"it was the season of light it was the season of darkness\n" +
		"it was the spring of hope it was the winter of despair"

	whitespace := regexp.MustCompile(`\s+`)
	text := whitespace.ReplaceAllString(tinyTale, " ")
	expected := "st of times it was the "
	assert.Equal(expected, Lrs(text))
}
