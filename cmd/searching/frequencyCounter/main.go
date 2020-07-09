package main

//    Read in a list of words from standard input and print out
//    the most frequently occurring word that has length greater than
//    a given threshold.
//
//    % go run main.go 1  < tinyTale.txt
//    it 10
//
//    % go run main.go  8 < tale.txt
//    business 122
//
//    % go run main.go 10 < leipzig1M.txt
//    government 24763

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/searching"
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
	distinct, wordscount := 0, 0
	minLen, _ := strconv.Atoi(os.Args[1])
	st := searching.NewBinarySearchST()
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		word := stdin.ReadString()
		if len(word) < minLen {
			continue
		}
		wordscount++
		if ok, _ := st.Contains(words(word)); !ok {
			//nolint:errcheck
			st.Put(words(word), 1)
			distinct++
		} else {
			val, _ := st.Get(words(word))
			//nolint:errcheck
			st.Put(words(word), val.(int)+1)
		}
	}

	// find a key with the highest frequency count
	max := words("")
	//nolint:errcheck
	st.Put(max, 0)
	for _, word := range st.Keys() {
		m, _ := st.Get(word)
		n, _ := st.Get(max)
		if m.(int) > n.(int) {
			max = word.(words)
		}
	}
	val, _ := st.Get(max)
	fmt.Println(max, "  ", val.(int))
	fmt.Println("distinct = ", distinct)
	fmt.Println("words   = ", wordscount)
}
