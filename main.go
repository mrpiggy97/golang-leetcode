package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mrpiggy97/golang-leetcode/arithmetic"
	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func main() {
	var initialString = "fabian#@!-_"
	var Sentence string = "this-is-a-sentence"
	var separatedSentence []string = strings.Split(Sentence, "-")
	var transformThis string = "christopher"
	var firstNumber int = 18
	var secondNumber int = 12
	initialString = stringMethods.ReverseString(initialString)
	fmt.Printf("%v %v\n", initialString, strings.ToUpper(initialString))
	fmt.Printf("%v\n", separatedSentence)
	fmt.Printf("%v\n", transformThis)
	arithmetic.DivisorFind(firstNumber)
	arithmetic.DivisorFind(secondNumber)
	arithmetic.DivisorFind(4500000000000000000)
	var newS string = stringMethods.CreateNewString("NyffsGeyylB")
	fmt.Printf("%v\n", newS)
	var otherString string = stringMethods.RepeatString(4, "a")
	fmt.Printf("%v\n", otherString)
	fmt.Printf("%v\n", arithmetic.MultiplyAllMembers([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v\n", arithmetic.DoubleTheAge(36, 7))
	fmt.Printf("%v\n", stringMethods.CompareEndOfString("", ""))
	fmt.Printf("%v\n", stringMethods.ChangeToCamelCase("the-stealth-warrior"))
	fmt.Printf("%v\n", arithmetic.Race(720, 850, 70))
	fmt.Printf("%v\n", stringMethods.CheckIfUpperCase("CCsMO LA MADRE"))
	fmt.Printf("%v\n", arithmetic.SumMembers(195))
	fmt.Printf("%v\n", stringMethods.DuplicateEncode("fAbian"))

	var names []string = []string{"fabian,", "jesus,", "rivas,"}
	var appendThis []string = []string{"chris,", "agrippa,", "augustus,"}
	var funcs []string = []string{"{", "}"}
	for _, member := range appendThis {
		names = stringMethods.InsertAtPosition(1, names, member)
		funcs = stringMethods.InsertAtPosition(1, funcs, member)
		fmt.Printf("%v\n", names)
		fmt.Printf("%v\n", funcs)
	}

	var name string = "fabian Is the bOSSS"
	fmt.Printf("%v\n", stringMethods.GetConvertedString(name, make(chan string, 1)))
	var check []byte = []byte("")
	fmt.Println(len(check))
	for _, byteVal := range check {
		fmt.Println(unicode.IsLetter(rune(byteVal)))
		fmt.Println(unicode.IsDigit(rune(byteVal)))
	}
}
