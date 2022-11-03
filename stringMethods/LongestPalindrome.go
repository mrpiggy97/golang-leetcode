package stringMethods

import (
	"strings"
)

func IsPalindrome(subStr []string) bool {
	var endIndex int = int(len(subStr) / 2)
	var currentTail int = len(subStr) - 1
	for index := range subStr[0:endIndex] {
		if subStr[index] != subStr[currentTail] {
			return false
		}
		currentTail = currentTail - 1
	}
	return true
}

func GetClosureIndexes(index int, subStr []string) []int {
	var closuresIndexes []int = []int{}
	for i, val := range subStr {
		if i != index {
			if val == subStr[index] {
				closuresIndexes = append(closuresIndexes, i)
			}
		}
	}
	return closuresIndexes
}

func GetSubStr(str []string, startingIndex int, endIndex int) []string {
	var newStr []string = []string{}
	for i := startingIndex; i <= endIndex; i++ {
		newStr = append(newStr, str[i])
	}
	return newStr
}

func LongestPalindrome(str string) string {
	var strSlice []string = strings.Split(str, "")
	var selectedSlice []string = []string{string(str[0])}
	for index := range strSlice {
		var closureIndexes []int = GetClosureIndexes(index, strSlice)
		if len(closureIndexes) > 0 {
			// we start with the largest possible palindrome
			for currentIndex := len(closureIndexes) - 1; currentIndex >= 0; currentIndex-- {
				var subSlice []string = GetSubStr(strSlice, index, closureIndexes[currentIndex])
				var subSliceIsPalindrome bool = IsPalindrome(subSlice)
				if subSliceIsPalindrome {
					if len(subSlice) > len(selectedSlice) {
						selectedSlice = subSlice
					}
				}
			}
		}
	}
	return strings.Join(selectedSlice, "")
}
