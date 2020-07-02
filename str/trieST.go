package str

import (
	"strings"

	"github.com/pkg/errors"
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
func (st *TrieST) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("argument to Get() is empty string")
	}
	x := st.get(st.root, key, 0)
	if x == nil {
		return nil, nil
	}
	return x.val, nil
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
func (st *TrieST) Contains(key string) (bool, error) {
	if key == "" {
		return false, errors.New("argument to Contains() is empty string")
	}
	val, _ := st.Get(key)
	return val != nil, nil
}

// Put inserts the key-value pair into the symbol table,
// overwriting the old value with the new value if the key is already in the symbol table.
func (st *TrieST) Put(key string, val interface{}) error {
	if key == "" {
		return errors.New("first argument to Put() is empty string")
	}
	if val == nil {
		//nolint:errcheck
		st.Delete(key)
		return nil
	}
	st.root = st.put(st.root, key, val, 0)
	return nil
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
	var s strings.Builder
	s.WriteString(prefix)
	st.collectPrefix(st.get(st.root, prefix, 0), &s, &results)
	return results
}

func (st *TrieST) collectPrefix(x *node, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}
	if x.val != nil {
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
	var s strings.Builder
	st.collectMatch(st.root, &s, pattern, &results)
	return results
}

func (st *TrieST) collectMatch(x *node, prefix *strings.Builder, pattern string, results *[]string) {
	if x == nil {
		return
	}
	d := prefix.Len()
	str := prefix.String()
	if d == len(pattern) && x.val != nil {
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

// LongestPrefixOf returns the string in the symbol table that
// is the longest prefix of query, or nil, if no such string.
func (st *TrieST) LongestPrefixOf(query string) (string, error) {
	if query == "" {
		return "", errors.New("argument to LongestPrefixOf() is empty string")
	}
	length := st.longestPrefixOf(st.root, query, 0, -1)
	if length == -1 {
		return "", nil
	}
	return query[:length], nil
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
func (st *TrieST) Delete(key string) error {
	if key == "" {
		return errors.New("argument to Delete() is empty string")
	}
	st.root = st.delete(st.root, key, 0)
	return nil
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
