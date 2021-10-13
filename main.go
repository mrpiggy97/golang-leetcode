package main

import (
	"fmt"
	"strings"

	"github.com/mrpiggy97/golang-leetcode/arithmetic"
	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func main() {
	var initialString = "fabian#@!-_"
	var Sentence string = "this-is-a-sentence"
	var separatedSentence []string = strings.Split(Sentence, "-")
	var convertToCamelCase string = "the-stealth-warrior"
	var transformThis string = "christopher"
	var firstNumber int = 18
	var secondNumber int = 12
	convertToCamelCase = stringMethods.ChangeToCamelCase(convertToCamelCase)
	initialString = stringMethods.ReverseString(initialString)
	fmt.Printf("%v %v\n", initialString, strings.ToUpper(initialString))
	fmt.Printf("%v\n", separatedSentence)
	fmt.Printf("%v\n", convertToCamelCase)
	fmt.Printf("%v\n", transformThis)
	arithmetic.DivisorFind(firstNumber)
	arithmetic.DivisorFind(secondNumber)
	arithmetic.DivisorFind(4500000000000000000)
	var newS string = stringMethods.CreateNewString("NyffsGeyylB")
	fmt.Printf("%v\n", newS)
}
