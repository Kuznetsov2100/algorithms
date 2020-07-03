package main

//   Symbol table with string keys, implemented using a ternary search
//    trie (TST).

//    % go run main.go < shellsST.txt
//    keys(""):
//    by 4
//    sea 6
//    sells 1
//    she 0
//    shells 3
//    shore 7
//    the 5

//    longestPrefixOf("shellsort"):
//    shells

//    keysWithPrefix("shor"):
//    shore

//    keysThatMatch(".he.l."):
//    shells

//    theory the now is the time for all good men

//    Remarks
//    --------
//      - can't use a key that is the empty string ""

import (
	"fmt"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	stdin := stdin.NewStdIn()
	st := str.NewTST()
	for i := 0; !stdin.IsEmpty(); i++ {
		key := stdin.ReadString()
		st.Put(key, i)
	}

	if st.Size() < 100 {
		fmt.Println("keys(\"\"):")
		for _, key := range st.Keys() {
			fmt.Println(key, " ", st.Get(key).(int))
		}
		fmt.Println()
	}

	fmt.Println("LongestPrefixOf(\"shellsort\"):")
	fmt.Println(st.LongestPrefixOf("shellsort"))
	fmt.Println()

	fmt.Println("LongestPrefixOf(\"shell\"):")
	fmt.Println(st.LongestPrefixOf("shell"))
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
