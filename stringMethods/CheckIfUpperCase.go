package stringMethods

import (
	"strings"
)

func CheckIfUpperCase(word string) bool {
	return strings.ToUpper(word) == word
}
