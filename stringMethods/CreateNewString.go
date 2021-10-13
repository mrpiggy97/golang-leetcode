package stringMethods

import (
	"strings"
)

func CreateNewString(stringInstance string) string {
	var newString string = ""
	for i := 0; i < len(stringInstance); i++ {
		for numberOfTimesAppended := 0; numberOfTimesAppended <= i; numberOfTimesAppended++ {
			if numberOfTimesAppended == 0 {
				newString = newString + strings.ToUpper(string(stringInstance[i]))
			} else {
				newString = newString + strings.ToLower(string(stringInstance[i]))
			}
		}

		if i < len(stringInstance)-1 {
			newString = newString + "-"
		}
	}

	return newString
}
