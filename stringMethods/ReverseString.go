package stringMethods

func ReverseString(initialString string) string {
	var modifiedString string = ""
	for i := len(initialString) - 1; i >= 0; i-- {
		modifiedString = modifiedString + string(initialString[i])
	}

	return modifiedString
}
