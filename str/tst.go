package str

import "strings"

// TST stuct represents an symbol table of key-value pairs, with string keys and generic values.
// It supports the usual put, get, contains, delete, size, and is-empty methods. It also provides
// character-based methods for finding the string in the symbol table that is the longest prefix
// of a given prefix, finding all strings in the symbol table that start with a given prefix,
// and finding all strings in the symbol table that match a given pattern. A symbol table implements
// the associative array abstraction: when associating a value with a key that is already in the symbol
// table, the convention is to replace the old value with the new value. Unlike Map, this struct uses
// the convention that values cannot be nil—setting the value associated with a key to nil
// is equivalent to deleting the key from the symbol table.
// This implementation uses a ternary search trie.
type TST struct {
	n    int    // size
	root *nodeC // root of TST
}

type nodeC struct {
	c     byte        // character
	left  *nodeC      // left subtries
	mid   *nodeC      // mid subtries
	right *nodeC      // right subtries
	val   interface{} // value associated with string
}

// NewTST initializes an empty string symbol table.
func NewTST() *TST {
	return &TST{}
}

// Size returns the number of key-value pairs in this symbol table.
func (st *TST) Size() int {
	return st.n
}

// IsEmpty returns true if the symbol table is empty.
func (st *TST) IsEmpty() bool {
	return st.n == 0
}

// Contains returns true if the symbol table contains the given key.
func (st *TST) Contains(key string) bool {
	return st.Get(key) != nil
}

// Get returns the value associated with the given key.
func (st *TST) Get(key string) interface{} {
	x := st.get(st.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func (st *TST) get(x *nodeC, key string, d int) *nodeC {
	if x == nil {
		return nil
	}
	if key == "" {
		panic("key must have length >= 1")
	}
	if c := key[d]; c < x.c {
		return st.get(x.left, key, d)
	} else if c > x.c {
		return st.get(x.right, key, d)
	} else if d < len(key)-1 {
		return st.get(x.mid, key, d+1)
	} else {
		return x
	}
}

// Put inserts the key-value pair into the symbol table,
// overwriting the old value with the new value if the key is already in the symbol table.
func (st *TST) Put(key string, val interface{}) {
	if !st.Contains(key) {
		st.n++
	} else if val == nil {
		st.n-- // delete existing key
	}
	st.root = st.put(st.root, key, val, 0)
}

func (st *TST) put(x *nodeC, key string, val interface{}, d int) *nodeC {
	c := key[d]
	if x == nil {
		x = &nodeC{c: c}
	}
	if c < x.c {
		x.left = st.put(x.left, key, val, d)
	} else if c > x.c {
		x.right = st.put(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.mid = st.put(x.mid, key, val, d+1)
	} else {
		x.val = val
	}
	return x
}

// LongestPrefixOf returns the string in the symbol table that
// is the longest prefix of query, or nil, if no such string.
func (st *TST) LongestPrefixOf(query string) string {
	if query == "" {
		return ""
	}
	length := 0
	x := st.root
	for i := 0; x != nil && i < len(query); {
		if c := query[i]; c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.val != nil {
				length = i
			}
			x = x.mid
		}
	}
	return query[:length]
}

// Keys returns all keys in the symbol table as a slice
func (st *TST) Keys() (keys []string) {
	var s strings.Builder
	st.collectPrefix(st.root, &s, &keys)
	return keys
}

// KeysWithPrefix returns all of the keys in the set that start with prefix.
func (st *TST) KeysWithPrefix(prefix string) (results []string) {
	x := st.get(st.root, prefix, 0)
	if x == nil {
		return results
	}
	if x.val != nil {
		results = append(results, prefix)
	}
	var s strings.Builder
	s.WriteString(prefix)
	st.collectPrefix(x.mid, &s, &results)
	return results
}

// in-order traversal gives nodes in non-decreasing order
func (st *TST) collectPrefix(x *nodeC, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}
	st.collectPrefix(x.left, prefix, results) // Step 1 - Recursively traverse left subtree.
	str := prefix.String()
	if x.val != nil {
		*results = append(*results, str+string(x.c))
	}
	prefix.WriteByte(x.c)
	st.collectPrefix(x.mid, prefix, results) // Step 2 - Visit mid node.
	prefix.Reset()
	prefix.WriteString(str)
	st.collectPrefix(x.right, prefix, results) // step 3 - Recursively traverse right subtree.
}

// KeysThatMatch returns all of the keys in the symbol table that match pattern,
// where "." symbol is treated as a wildcard character.
func (st *TST) KeysThatMatch(pattern string) (results []string) {
	var s strings.Builder
	st.collectMatch(st.root, &s, 0, pattern, &results)
	return results
}

func (st *TST) collectMatch(x *nodeC, prefix *strings.Builder, i int, pattern string, results *[]string) {
	if x == nil {
		return
	}
	c := pattern[i]
	str := prefix.String()
	if string(c) == "." || c < x.c {
		st.collectMatch(x.left, prefix, i, pattern, results)
	}
	if string(c) == "." || c == x.c {
		if i == len(pattern)-1 && x.val != nil {
			*results = append(*results, str+string(x.c))
		}
		if i < len(pattern)-1 {
			prefix.WriteByte(x.c)
			st.collectMatch(x.mid, prefix, i+1, pattern, results)
			prefix.Reset()
			prefix.WriteString(str)
		}
	}
	if string(c) == "." || c > x.c {
		st.collectMatch(x.right, prefix, i, pattern, results)
	}
}
