package searching

import (
	"errors"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

type ST struct {
	st *treemap.Map
}

func NewST(comparator utils.Comparator) *ST {
	return &ST{st: treemap.NewWith(comparator)}
}

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

func (s *ST) Remove(key Key) error {
	if key == nil {
		return errors.New("calls Remove() with nil key")
	}
	s.st.Remove(key)
	return nil
}

func (s *ST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("calls Contains() with nil key")
	}
	_, ok := s.st.Get(key)
	return ok, nil
}

func (s *ST) Size() int {
	return s.st.Size()
}

func (s *ST) IsEmpty() bool {
	return s.st.Empty()
}

func (s *ST) Keys() (k []Key) {
	for _, val := range s.st.Keys() {
		k = append(k, val.(Key))
	}
	return k
}

func (s *ST) Min() (Key, error) {
	if s.IsEmpty() {
		return nil, errors.New("calls Min() with empty symbol table")
	}
	k, _ := s.st.Min()
	return k.(Key), nil
}

func (s *ST) Max() (Key, error) {
	if s.IsEmpty() {
		return nil, errors.New("calls Max() with empty symbol table")
	}
	k, _ := s.st.Max()
	return k.(Key), nil
}

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
