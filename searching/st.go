package searching

import (
	"github.com/pkg/errors"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

// ST struct represents an ordered symbol table of generic key-value pairs.
// It requires that the key type implements the Key interface.
type ST struct {
	st *treemap.Map
}

// NewST initializes an empty symbol table.
func NewST(comparator utils.Comparator) *ST {
	return &ST{st: treemap.NewWith(comparator)}
}

// Get returns the value associated with the given key in this symbol table.
func (s *ST) Get(key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("calls Get() wit nil key")
	}
	if val, ok := s.st.Get(key); ok {
		return val.(Value), nil
	} else {
		return nil, nil
	}
}

// Put inserts the specified key-value pair into the symbol table, overwriting the old value
// with the new value if the symbol table already contains the specified key.
func (s *ST) Put(key Key, val Value) error {
	if key == nil {
		return errors.New("calls Put() with nil key")
	}
	if val == nil {
		s.st.Remove(key)
		return nil
	}
	s.st.Put(key, val)
	return nil
}

// Remove removes the specified key and its associated value from this symbol table
// (if the key is in this symbol table).
func (s *ST) Remove(key Key) error {
	if key == nil {
		return errors.New("calls Remove() with nil key")
	}
	s.st.Remove(key)
	return nil
}

// Contains returns true if this symbol table contain the given key.
func (s *ST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("calls Contains() with nil key")
	}
	_, ok := s.st.Get(key)
	return ok, nil
}

// Size returns the number of key-value pairs in this symbol table.
func (s *ST) Size() int {
	return s.st.Size()
}

// IsEmpty returns true if this symbol table is empty.
func (s *ST) IsEmpty() bool {
	return s.st.Empty()
}

// Keys returns all keys in this symbol table.
func (s *ST) Keys() (k []Key) {
	for _, val := range s.st.Keys() {
		k = append(k, val.(Key))
	}
	return k
}

// Min returns all keys in this symbol table.
func (s *ST) Min() (Key, error) {
	if s.IsEmpty() {
		return nil, errors.New("calls Min() with empty symbol table")
	}
	k, _ := s.st.Min()
	return k.(Key), nil
}

// Max returns the largest key in this symbol table.
func (s *ST) Max() (Key, error) {
	if s.IsEmpty() {
		return nil, errors.New("calls Max() with empty symbol table")
	}
	k, _ := s.st.Max()
	return k.(Key), nil
}

// Ceiling returns the smallest key in this symbol table greater than or equal to key.
func (s *ST) Ceiling(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to ceiling() with nil key")
	}
	k, _ := s.st.Ceiling(key)
	if k == nil {
		return nil, errors.New("argument to Ceiling() is too large")
	}
	return k.(Key), nil
}

// Floor returns the largest key in this symbol table less than or equal to key.
func (s *ST) Floor(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Floor() is nil key")
	}
	k, _ := s.st.Floor(key)
	if k == nil {
		return nil, errors.New("argument to Floor() is too small")
	}
	return k.(Key), nil
}
