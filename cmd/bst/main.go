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
	distinct, wordscount := 0, 0
	minLen, _ := strconv.Atoi(os.Args[1])
	bst := searching.NewBST()
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		word := stdin.ReadString()
		if len(word) < minLen {
			continue
		}
		wordscount++
		if ok, _ := bst.Contains(words(word)); !ok {
			//nolint:errcheck
			bst.Put(words(word), 1)
			distinct++
		} else {
			//nolint:errcheck
			bst.Put(words(word), bst.Get(words(word)).(int)+1)
		}
	}

	// find a key with the highest frequency count
	max := words("")
	//nolint:errcheck
	bst.Put(max, 0)
	for _, word := range bst.Keys() {
		if bst.Get(word).(int) > bst.Get(max).(int) {
			max = word.(words)
		}
	}
	fmt.Println(max, "  ", bst.Get(max).(int))
	fmt.Println("distinct = ", distinct)
	fmt.Println("words   = ", wordscount)
}
