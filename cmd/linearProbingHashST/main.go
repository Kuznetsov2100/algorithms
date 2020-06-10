package main

//    Read in a list of stringHashKey from standard input and print out
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

type stringHashKey = searching.StringHashKey

func main() {
	distinct, wordscount := 0, 0
	minLen, _ := strconv.Atoi(os.Args[1])
	bst := searching.NewSeparateChainingHashST(1667)
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		word := stdin.ReadString()
		if len(word) < minLen {
			continue
		}
		wordscount++
		if ok, _ := bst.Contains(stringHashKey(word)); !ok {
			//nolint:errcheck
			bst.Put(stringHashKey(word), 1)
			distinct++
		} else {

			val, _ := bst.Get(stringHashKey(word))
			//nolint:errcheck
			bst.Put(stringHashKey(word), val.(int)+1)
		}
	}

	// find a key with the highest frequency count
	max := stringHashKey("")
	//nolint:errcheck
	bst.Put(max, 0)
	for _, word := range bst.Keys() {
		m, _ := bst.Get(word)
		n, _ := bst.Get(max)
		if m.(int) > n.(int) {
			max = word.(stringHashKey)
		}
	}
	maxfrequency, _ := bst.Get(max)
	fmt.Println(max, "  ", maxfrequency.(int))
	fmt.Println("distinct = ", distinct)
	fmt.Println("words = ", wordscount)
}
