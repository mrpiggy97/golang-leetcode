package stringMethods

import "strings"

// Converts roman integer into regular int
func RomanToInt(roman string) int {
	var romanConverstion map[string]int = make(map[string]int)
	romanConverstion["I"] = 1
	romanConverstion["V"] = 5
	romanConverstion["X"] = 10
	romanConverstion["L"] = 50
	romanConverstion["C"] = 100
	romanConverstion["D"] = 500
	romanConverstion["M"] = 1000
	var allowedSubstractions map[string]string = make(map[string]string)
	allowedSubstractions["V"] = "I"
	allowedSubstractions["X"] = "I"
	allowedSubstractions["L"] = "X"
	allowedSubstractions["C"] = "X"
	allowedSubstractions["D"] = "C"
	allowedSubstractions["M"] = "C"
	var previousChar string = ""
	var total int = 0
	var romanSlice []string = strings.Split(roman, "")
	for index, value := range romanSlice {
		var currentConverstion int = romanConverstion[value]
		if index == 0 {
			previousChar = value
			total += currentConverstion
		} else {
			var valueToSubstract string = allowedSubstractions[value]
			if valueToSubstract == previousChar {
				var previousConverstion int = romanConverstion[previousChar]
				total -= previousConverstion
				var secondTotal int = currentConverstion - previousConverstion
				total += secondTotal
			} else {
				total += currentConverstion
			}
			previousChar = value
		}
	}
	return total
}
