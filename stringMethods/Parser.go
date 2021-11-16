package stringMethods

import "unicode"

func Parser(stringInstance string) bool {
	if len(stringInstance) == 0 {
		return false
	}
	var byteSlice []byte = []byte(stringInstance)
	for _, byteCharacter := range byteSlice {
		var isDigit bool = unicode.IsDigit(rune(byteCharacter))
		var isLetter bool = unicode.IsLetter(rune(byteCharacter))
		if !isDigit && !isLetter {
			return false
		}
	}

	return true
}
