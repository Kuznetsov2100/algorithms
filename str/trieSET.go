package str

import (
	"strings"
)

// TrieSET struct represents an ordered set of strings over the extended ASCII alphabet.
// It supports the usual add, contains, and delete methods. It also provides character-based methods
// for finding the string in the set that is the longest prefix of a given prefix,
// finding all strings in the set that start with a given prefix, and finding all strings
// in the set that match a given pattern.
// This implementation uses a 256-way trie. The add, contains, delete, and longest prefix methods
// take time proportional to the length of the key (in the worst case). Construction takes constant time.
type TrieSET struct {
	root *nodeB // root of trie
	n    int    // number of keys in trie
}

// R-way trie nodeB
type nodeB struct {
	isString bool
	next     [asciiR]*nodeB
}

// NewTrieSET initializes an empty set of strings.
func NewTrieSET() *TrieSET {
	return &TrieSET{}
}

// Size returns the number of strings in the set.
func (st *TrieSET) Size() int {
	return st.n
}

// IsEmpty returns true if the set is empty.
func (st *TrieSET) IsEmpty() bool {
	return st.n == 0
}

func (st *TrieSET) get(x *nodeB, key string, d int) *nodeB {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	return st.get(x.next[key[d]], key, d+1)
}

// Contains returns true if the set contains the given key.
func (st *TrieSET) Contains(key string) bool {
	x := st.get(st.root, key, 0)
	if x == nil {
		return false
	}
	return x.isString
}

// Add adds the key to the set if it is not already present.
func (st *TrieSET) Add(key string) {
	st.root = st.add(st.root, key, 0)
}

func (st *TrieSET) add(x *nodeB, key string, d int) *nodeB {
	if x == nil {
		x = &nodeB{}
	}
	if d == len(key) {
		if !x.isString {
			st.n++
		}
		x.isString = true
	} else {
		x.next[key[d]] = st.add(x.next[key[d]], key, d+1)
	}
	return x
}

// Iterator returns all of the keys in the set as a slice
func (st *TrieSET) Iterator() []string {
	return st.KeysWithPrefix("")
}

// KeysWithPrefix returns all of the keys in the set that start with prefix.
func (st *TrieSET) KeysWithPrefix(prefix string) (results []string) {
	var s strings.Builder
	s.WriteString(prefix)
	st.collectPrefix(st.get(st.root, prefix, 0), &s, &results)
	return results
}

func (st *TrieSET) collectPrefix(x *nodeB, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}
	if x.isString {
		*results = append(*results, prefix.String())
	}
	str := prefix.String()
	for c := 0; c < asciiR; c++ {
		prefix.Reset()
		prefix.WriteString(str)
		prefix.WriteByte(byte(c))
		st.collectPrefix(x.next[c], prefix, results)
	}
}

// KeysThatMatch returns all of the keys in the set that match the pattern,
// where "." symbol is treated as a wildcard character.
func (st *TrieSET) KeysThatMatch(pattern string) (results []string) {
	var s strings.Builder
	st.collectMatch(st.root, &s, pattern, &results)
	return results
}

func (st *TrieSET) collectMatch(x *nodeB, prefix *strings.Builder, pattern string, results *[]string) {
	if x == nil {
		return
	}
	d := prefix.Len()
	str := prefix.String()
	if d == len(pattern) && x.isString {
		*results = append(*results, str)
	}
	if d == len(pattern) {
		return
	}
	c := string(pattern[d])
	if c == "." {
		for ch := 0; ch < asciiR; ch++ {
			prefix.Reset()
			prefix.WriteString(str)
			prefix.WriteByte(byte(ch))
			st.collectMatch(x.next[ch], prefix, pattern, results)
		}
	} else {
		prefix.WriteString(c)
		st.collectMatch(x.next[c[0]], prefix, pattern, results)
	}
}

// LongestPrefixOf returns the string in the set that
// is the longest prefix of query, or nil, if no such string.
func (st *TrieSET) LongestPrefixOf(query string) string {
	length := st.longestPrefixOf(st.root, query, 0, -1)
	if length == -1 {
		return ""
	}
	return query[:length]
}

func (st *TrieSET) longestPrefixOf(x *nodeB, query string, d, length int) int {
	if x == nil {
		return length
	}
	if x.isString {
		length = d
	}
	if d == len(query) {
		return length
	}
	return st.longestPrefixOf(x.next[query[d]], query, d+1, length)
}

// Delete removes the key from the set if the key is present.
func (st *TrieSET) Delete(key string) {
	st.root = st.delete(st.root, key, 0)
}

func (st *TrieSET) delete(x *nodeB, key string, d int) *nodeB {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.isString {
			st.n--
		}
		x.isString = false
	} else {
		x.next[key[d]] = st.delete(x.next[key[d]], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.isString {
		return x
	}
	for c := 0; c < asciiR; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}
