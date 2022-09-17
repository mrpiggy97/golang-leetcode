package stringMethods

import (
	"strings"
)

func FirstUniqueChar(s string) int {
	var slice []string = strings.Split(s, "")
	var count = 0
	for index, str := range slice {
		count = strings.Count(s, str)
		if count == 1 {
			return index
		}
	}
	return -1
}
