package str

import (
	"bytes"
)

// TrieST struct represents an symbol table of key-value pairs, with string keys and generic values.
// It supports the usual put, get, contains, delete, size, and is-empty methods. It also provides
// character-based methods for finding the string in the symbol table that is the longest prefix of
// a given prefix, finding all strings in the symbol table that start with a given prefix, and finding
// all strings in the symbol table that match a given pattern. A symbol table implements the associative
// array abstraction: when associating a value with a key that is already in the symbol table,
// the convention is to replace the old value with the new value. Unlike Map,
// this struct uses the convention that values cannot be nil—setting the value associated with a key
// to nil is equivalent to deleting the key from the symbol table.
// This implementation uses a 256-way trie. The put, contains, delete, and longest prefix operations
// take time proportional to the length of the key (in the worst case). Construction takes constant time.
// The size, and is-empty operations take constant time. Construction takes constant time.
type TrieST struct {
	root *node // root of trie
	n    int   // number of keys in trie
}

// R-way trie node
type node struct {
	val  interface{}
	next [asciiR]*node
}

// NewTrieST initializes an empty string symbol table.
func NewTrieST() *TrieST {
	return &TrieST{}
}

// Size returns the number of key-value pairs in this symbol table.
func (st *TrieST) Size() int {
	return st.n
}

// IsEmpty returns true if the symbol table is empty.
func (st *TrieST) IsEmpty() bool {
	return st.n == 0
}

// Get returns the value associated with the given key.
func (st *TrieST) Get(key string) interface{} {

	x := st.get(st.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func (st *TrieST) get(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	return st.get(x.next[key[d]], key, d+1)
}

// Contains returns true if the symbol table contains the given key.
func (st *TrieST) Contains(key string) bool {
	return st.Get(key) != nil
}

// Put inserts the key-value pair into the symbol table,
// overwriting the old value with the new value if the key is already in the symbol table.
func (st *TrieST) Put(key string, val interface{}) {
	if val == nil {
		st.Delete(key)
	}
	st.root = st.put(st.root, key, val, 0)
}

func (st *TrieST) put(x *node, key string, val interface{}, d int) *node {
	if x == nil {
		x = &node{}
	}
	if d == len(key) {
		if x.val == nil {
			st.n++
		}
		x.val = val
		return x
	}
	x.next[key[d]] = st.put(x.next[key[d]], key, val, d+1)
	return x
}

// Keys returns all keys in the symbol table as a slice
func (st *TrieST) Keys() []string {
	return st.KeysWithPrefix("")
}

// KeysWithPrefix returns all of the keys in the set that start with prefix.
func (st *TrieST) KeysWithPrefix(prefix string) (results []string) {
	var b bytes.Buffer
	b.WriteString(prefix)
	st.collectPrefix(st.get(st.root, prefix, 0), &b, &results)
	return results
}

func (st *TrieST) collectPrefix(x *node, prefix *bytes.Buffer, results *[]string) {
	if x == nil {
		return
	}
	if x.val != nil {
		*results = append(*results, prefix.String())
	}
	for c := 0; c < asciiR; c++ {
		prefix.WriteByte(byte(c))
		st.collectPrefix(x.next[c], prefix, results)
		prefix.Truncate(prefix.Len() - 1)
	}
}

// this is the equivalent version of KeysWithPrefix, use string instead of strings.builder,speed is slower
/*func (st *TrieST) KeysWithPrefix(prefix string) (results []string) {
	st.collectPrefix(st.get(st.root, prefix, 0), prefix, &results)
	return results
}

func (st *TrieST) collectPrefix(x *node, prefix string, results *[]string) {
	if x == nil {
		return
	}
	if x.val != nil {
		*results = append(*results, prefix)
	}
	for c := 0; c < asciiR; c++ {
		prefix = prefix + string(c)
		st.collectPrefix(x.next[c], prefix, results)
		prefix = prefix[:len(prefix)-1]
	}
}*/

// KeysThatMatch returns all of the keys in the symbol table that match pattern,
// where "." symbol is treated as a wildcard character.
func (st *TrieST) KeysThatMatch(pattern string) (results []string) {
	var b bytes.Buffer
	st.collectMatch(st.root, &b, pattern, &results)
	return results
}

func (st *TrieST) collectMatch(x *node, prefix *bytes.Buffer, pattern string, results *[]string) {
	if x == nil {
		return
	}
	d := prefix.Len()
	if d == len(pattern) && x.val != nil {
		*results = append(*results, prefix.String())
	}
	if d == len(pattern) {
		return
	}
	if c := pattern[d]; c == '.' {
		for ch := 0; ch < asciiR; ch++ {
			prefix.WriteByte(byte(ch))
			st.collectMatch(x.next[ch], prefix, pattern, results)
			prefix.Truncate(prefix.Len() - 1)
		}
	} else {
		prefix.WriteByte(c)
		st.collectMatch(x.next[c], prefix, pattern, results)
		prefix.Truncate(prefix.Len() - 1)
	}
}

// LongestPrefixOf returns the string in the symbol table that
// is the longest prefix of query, or nil, if no such string.
func (st *TrieST) LongestPrefixOf(query string) string {
	length := st.longestPrefixOf(st.root, query, 0, -1)
	if length == -1 {
		return ""
	}
	return query[:length]
}

func (st *TrieST) longestPrefixOf(x *node, query string, d, length int) int {
	if x == nil {
		return length
	}
	if x.val != nil {
		length = d
	}
	if d == len(query) {
		return length
	}
	return st.longestPrefixOf(x.next[query[d]], query, d+1, length)
}

// Delete removes the key from the set if the key is present.
func (st *TrieST) Delete(key string) {
	st.root = st.delete(st.root, key, 0)
}

func (st *TrieST) delete(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.val != nil {
			st.n--
		}
		x.val = nil
	} else {
		x.next[key[d]] = st.delete(x.next[key[d]], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.val != nil {
		return x
	}
	for c := 0; c < asciiR; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}
