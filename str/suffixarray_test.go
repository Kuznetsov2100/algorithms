package str

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixArray(t *testing.T) {
	assert := assert.New(t)

	s := "ABRACADABRA!"
	suffix := NewSuffixArray(s)
	var out strings.Builder

	for i := 0; i < len(s); i++ {
		index := suffix.Index(i)
		ith := "\"" + s[index:min(index+50, len(s))] + "\""
		rank := suffix.Rank(s[index:])
		if i == 0 {
			fmt.Fprintf(&out, "%3d %3d %3s %3d %s\n", i, index, "-", rank, ith)
		} else {
			lcp := suffix.Lcp(i)
			fmt.Fprintf(&out, "%3d %3d %3d %3d %s\n", i, index, lcp, rank, ith)
		}
	}

	expected := "  0  11   -   0 \"!\"\n" +
		"  1  10   0   1 \"A!\"\n" +
		"  2   7   1   2 \"ABRA!\"\n" +
		"  3   0   4   3 \"ABRACADABRA!\"\n" +
		"  4   3   1   4 \"ACADABRA!\"\n" +
		"  5   5   1   5 \"ADABRA!\"\n" +
		"  6   8   0   6 \"BRA!\"\n" +
		"  7   1   3   7 \"BRACADABRA!\"\n" +
		"  8   4   0   8 \"CADABRA!\"\n" +
		"  9   6   0   9 \"DABRA!\"\n" +
		" 10   9   0  10 \"RA!\"\n" +
		" 11   2   2  11 \"RACADABRA!\"\n"

	assert.Equal(expected, out.String())
	assert.Equal(len(s), suffix.Length())
	assert.PanicsWithValue("invalid i", func() { suffix.Index(30) })
	assert.PanicsWithValue("invalid i", func() { suffix.Lcp(30) })
	assert.PanicsWithValue("invalid i", func() { suffix.Select(30) })
	assert.Equal("!", suffix.Select(0))
	assert.Equal(3, suffix.lcpSuffix("abc", "abc"))

	suffix1 := NewSuffixArray("")
	assert.Equal(0, suffix1.Rank("d"))
}
