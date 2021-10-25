package stringMethods

import (
	"strings"
)

func DuplicateEncode(stringInstance string) string {
	var newString string = ""
	stringInstance = strings.ToLower(stringInstance)
	for _, letter := range strings.Split(stringInstance, "") {
		if strings.Count(stringInstance, letter) > 1 {
			newString = newString + ")"
		} else {
			newString = newString + "("
		}
	}
	return newString
}
