package stringMethods

import "strings"

func LengthOfLastWord(s string) int {
	var slice []string = strings.Split(s, " ")
	var lastIndex int = len(slice) - 1
	var selectedString string = ""
	for index := lastIndex; index >= 0; index-- {
		var currentString string = slice[index]
		if len(currentString) > 0 {
			selectedString = currentString
		}
	}
	return len(selectedString)
}
