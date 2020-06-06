package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/searching"
	"github.com/handane123/algorithms/stdin"
)

type words string

func (w words) CompareTo(k searching.Key) int {
	t := k.(words)
	if w < t {
		return -1
	} else if w > t {
		return 1
	} else {
		return 0
	}
}

func main() {
	minLen, _ := strconv.Atoi(os.Args[1])
	st := searching.NewBinarySearchST()
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		word := stdin.ReadString()
		if len(word) < minLen {
			continue
		}
		wordKey := words(word)
		if ok, _ := st.Contains(wordKey); !ok {
			//nolint:errcheck
			st.Put(wordKey, 1)
		} else {
			val, _ := st.Get(wordKey)
			//nolint:errcheck
			st.Put(wordKey, val.(int)+1)
		}
	}

	max := words("")
	//nolint:errcheck
	st.Put(max, 0)
	for _, word := range st.Keys() {
		m, _ := st.Get(word)
		n, _ := st.Get(max)
		if m == nil {
			m = 0
		}
		if n == nil {
			n = 0
		}
		if m.(int) > n.(int) {
			max = word.(words)
		}
	}
	val, _ := st.Get(max)
	fmt.Printf("%s %d\n", max, val.(int))
}
