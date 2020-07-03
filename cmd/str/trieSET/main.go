package main

// An set for extended ASCII strings, implemented  using a 256-way trie.

//    % go run main.go < shellsST.txt
//    by
//    sea
//    sells
//    she
//    shells
//    shore
//    the

import (
	"fmt"

	"github.com/handane123/algorithms/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	stdin := stdin.NewStdIn()
	set := str.NewTrieSET()
	for !stdin.IsEmpty() {
		key := stdin.ReadString()
		set.Add(key)
	}

	if set.Size() < 100 {
		fmt.Println("keys(\"\"):")
		for _, key := range set.Iterator() {
			fmt.Println(key)
		}
		fmt.Println()
	}

	fmt.Println("LongestPrefixOf(\"shellsort\"):")
	fmt.Println(set.LongestPrefixOf("shellsort"))
	fmt.Println()

	fmt.Println("LongestPrefixOf(\"xshellsort\"):")
	fmt.Println(set.LongestPrefixOf("xshellsort"))
	fmt.Println()

	fmt.Println("LongestPrefixOf(\"shortening\"):")
	fmt.Println(set.LongestPrefixOf("shortening"))
	fmt.Println()

	fmt.Println("KeysWithPrefix(\"shor\"):")
	for _, s := range set.KeysWithPrefix("shor") {
		fmt.Println(s)
	}
	fmt.Println()

	fmt.Println("KeysThatMatch(\".he.l.\"):")
	for _, s := range set.KeysThatMatch(".he.l.") {
		fmt.Println(s)
	}

}
