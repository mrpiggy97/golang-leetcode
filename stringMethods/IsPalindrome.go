package stringMethods

import (
	"unicode"
)

func isValid(character rune) bool {
	var isDigit bool = unicode.IsDigit(character)
	var isLetter bool = unicode.IsLetter(character)
	if isDigit || isLetter {
		return true
	} else {
		return false
	}
}

func IsPalindromeLeetcode(s string) bool {
	var begginingIndex int = 0
	var endIndex int = len(s) - 1
	var currentValue rune
	var currentEndValue rune
	for begginingIndex < endIndex {
		var firstByte byte = s[begginingIndex]
		currentValue = rune(firstByte)
		var secondByte byte = s[endIndex]
		currentEndValue = rune(secondByte)
		var currentValueIsValid bool = isValid(currentValue)
		var currentEndValueIsValid bool = isValid(currentEndValue)
		if currentValueIsValid && currentEndValueIsValid {
			if !unicode.IsLower(currentValue) {
				currentValue = unicode.ToLower(currentValue)
			}
			if !unicode.IsLower(currentEndValue) {
				currentEndValue = unicode.ToLower(currentEndValue)
			}
			if currentValue != currentEndValue {
				return false
			} else {
				begginingIndex += 1
				endIndex -= 1
			}
		} else {
			if !currentValueIsValid {
				begginingIndex += 1
			}
			if !currentEndValueIsValid {
				endIndex -= 1
			}
		}
	}
	return true
}
