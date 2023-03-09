package stringMethods

func LengthOfLongestSubString(str string) int {
	if len(str) == 0 {
		return 0
	}
	var checker map[byte]int = make(map[byte]int)
	var index int = 0
	var maxLength int = 0
	for index < len(str) {
		var value byte = str[index]
		position, exists := checker[value]
		if exists {
			if index > maxLength {
				maxLength = index
			}
			if position+1 < len(str) {
				str = str[position+1:]
			}
			checker = make(map[byte]int)
			index = 0
		} else {
			checker[value] = index
			index += 1
		}
	}
	if len(str) > maxLength {
		maxLength = len(str)
	}
	return maxLength
}
