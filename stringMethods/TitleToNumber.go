package stringMethods

import (
	"math"
	"strings"
)

func GetStart(columnTitle string) (currentStart, currentEnd int, finalString string) {
	if len(columnTitle) == 1 {
		finalString = "A"
		return 1, 27, finalString
	}
	if len(columnTitle) == 2 {
		finalString = "AA"
		return 27, 27 + int(math.Pow(26, 2)), finalString
	}
	var currentPower float64 = 2
	currentStart = 27
	currentEnd = 27
	finalString = ""
	for currentPower < float64(len(columnTitle)) {
		currentStart = currentStart + int(math.Pow(26, currentPower))
		currentPower += 1
	}
	for len(finalString) < len(columnTitle) {
		finalString = finalString + "A"
	}
	currentEnd = currentStart + int(math.Pow(26, currentPower))
	return currentStart, currentEnd, finalString
}

// Algorithm goes as follows number = number + 26**len(str)-1 if letter
// does not coincide
func TitleToNumber(columnTitle string) int {
	var order [26]string = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	currentNumber, _, finalString := GetStart(columnTitle)
	var finalStringSlice []string = strings.Split(finalString, "")
	var columnTitleSlice []string = strings.Split(columnTitle, "")
	var indexInFinalStringSlice int = 0
	var indexInOrder int = 0
	var currentPower float64 = float64(len(finalString) - 1)
	for indexInFinalStringSlice < len(finalStringSlice) {
		if finalStringSlice[indexInFinalStringSlice] != columnTitleSlice[indexInFinalStringSlice] {
			currentNumber = currentNumber + int(math.Pow(26, currentPower))
			indexInOrder += 1
			finalStringSlice[indexInFinalStringSlice] = order[indexInOrder]
		} else {
			indexInFinalStringSlice += 1
			indexInOrder = 0
			currentPower -= 1
		}
	}
	return currentNumber
}
