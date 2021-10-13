package stringMethods

import (
	"fmt"
	"strings"
)

func ReverseString(initialString string) string {
	var modifiedString string = ""
	for i := len(initialString) - 1; i >= 0; i-- {
		modifiedString = modifiedString + string(initialString[i])
	}
	var uppercase string = strings.ToUpper(initialString)
	var byteUppercase []byte = []byte(uppercase)
	for index, character := range []byte(initialString) {
		fmt.Printf("%v %v\n", character, string(character))
		fmt.Printf("%v %v\n", byteUppercase[index], string(byteUppercase[index]))
	}
	return modifiedString
}
