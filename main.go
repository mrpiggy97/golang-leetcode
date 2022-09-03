package main

import (
	"fmt"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func main() {
	var i []int = []int{1, 3, 5, 6, 0, 0}
	var t []int = []int{4, 7}
	arrays.Merge(i, len(i), t, len(t))
	fmt.Println(i)
}
