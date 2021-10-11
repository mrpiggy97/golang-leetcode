package main

import (
	"fmt"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func main() {
	var initialString = "fabian"
	initialString = stringMethods.ReverseString(initialString)
	fmt.Printf("%v\n", initialString)
}
