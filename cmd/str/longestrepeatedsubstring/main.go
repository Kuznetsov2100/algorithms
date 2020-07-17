package main

import (
	"fmt"
	"regexp"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	space := regexp.MustCompile(`\s+`)
	text := space.ReplaceAllString(stdin.ReadAll(), " ")
	fmt.Printf("'%s'\n", str.Lrs(text))

}
