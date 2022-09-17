package stringMethods

import "strings"

func IsAnagram(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}
	var firstSlice []string = strings.Split(first, "")
	var secondSlice []string = strings.Split(second, "")
	var firstMap map[string]int = make(map[string]int)

	// fill maps
	for _, letter := range firstSlice {
		_, exists := firstMap[letter]
		if !exists {
			firstMap[letter] = strings.Count(first, letter)
		}
	}

	for _, letter := range secondSlice {
		countInFirstMap, existsInFirstMap := firstMap[letter]
		if !existsInFirstMap {
			return false
		}
		var countInSecondString int = strings.Count(second, letter)
		if countInFirstMap != countInSecondString {
			return false
		}

	}
	return true
}
