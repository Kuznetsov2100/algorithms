package main

//   A string symbol table for extended ASCII strings, implemented
//    using a 256-way trie.

//    % go run main.go < shellsST.txt
//    by 4
//    sea 6
//    sells 1
//    she 0
//    shells 3
//    shore 7
//    the 5

import (
	"fmt"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	stdin := stdin.NewStdIn()
	st := str.NewTrieST()
	for i := 0; !stdin.IsEmpty(); i++ {
		key := stdin.ReadString()
		//nolint:errcheck
		st.Put(key, i)
	}

	if st.Size() < 100 {
		fmt.Println("keys(\"\"):")
		for _, key := range st.Keys() {
			val, err := st.Get(key)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
			fmt.Println(key, " ", val.(int))
		}
		fmt.Println()
	}

	fmt.Println("LongestPrefixOf(\"shellsort\"):")
	fmt.Println(st.LongestPrefixOf("shellsort"))
	fmt.Println()

	fmt.Println("LongestPrefixOf(\"quicksort\"):")
	fmt.Println(st.LongestPrefixOf("quicksort"))
	fmt.Println()

	fmt.Println("KeysWithPrefix(\"shor\"):")
	for _, s := range st.KeysWithPrefix("shor") {
		fmt.Println(s)
	}
	fmt.Println()

	fmt.Println("KeysThatMatch(\".he.l.\"):")
	for _, s := range st.KeysThatMatch(".he.l.") {
		fmt.Println(s)
	}

}
