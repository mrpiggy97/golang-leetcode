package stringMethods

func CompareEndOfString(baseString string, stringToCompare string) bool {
	if len(stringToCompare) > len(baseString) {
		return false
	}
	var indexToStop int = len(baseString) - len(stringToCompare)
	var copy []byte = []byte(baseString[0:indexToStop])
	copy = append(copy, []byte(stringToCompare)...)
	return string(copy) == baseString
}
