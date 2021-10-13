package stringMethods

import "strings"

func IsUpper(character byte) bool {
	//we assume character is lower case
	var upperCaseUnicodeValue byte = byte(int(character) - 32)
	if string(upperCaseUnicodeValue) == strings.ToUpper(string(character)) {
		return false
	} else {
		return true
	}
}

func ChangeToCamelCase(stringInstance string) string {
	var stringSlice []string = []string{}
	var finalString string = ""
	if strings.Contains(stringInstance, "-") {
		stringSlice = strings.Split(stringInstance, "-")
	} else if strings.Contains(stringInstance, "_") {
		stringSlice = strings.Split(stringInstance, "_")
	}

	for index, word := range stringSlice {
		var byteSlice []byte = []byte(word)
		var firstByteIsUpper bool = IsUpper(byteSlice[0])
		if !firstByteIsUpper {
			byteSlice[0] = byte(int(byteSlice[0]) + 32)
		}
		stringSlice[index] = string(byteSlice)
		finalString = finalString + stringSlice[index]
	}

	return finalString
}
