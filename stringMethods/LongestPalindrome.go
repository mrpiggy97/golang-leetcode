package stringMethods

func IsPalindrome(str string, startingIndex, tailIndex int) bool {
	var subStr string = str[startingIndex:tailIndex] + string(str[tailIndex])
	var splitIndex int = int(len(subStr) / 2)
	var currentEnd int = len(subStr) - 1
	for i := 0; i < splitIndex; i++ {
		var head byte = subStr[i]
		var tail byte = subStr[currentEnd]
		if head != tail {
			return false
		}
		currentEnd = currentEnd - 1
	}
	return true
}

func LongestPalindrome(str string) string {
	if len(str) == 0 {
		return str
	}
	var selectedPalindrome string = string(str[0])
	for index, value := range str {
		for tail := len(str) - 1; tail > index; tail-- {
			var tailValue string = string(str[tail])
			if tailValue == string(value) {
				var length int = (tail - index) + 1
				if length < len(selectedPalindrome) {
					break
				}
				var isPalindrome bool = IsPalindrome(str, index, tail)
				if isPalindrome {
					var newStr = string(str[index:tail])
					newStr = newStr + tailValue
					if len(newStr) > len(selectedPalindrome) {
						selectedPalindrome = newStr
						break
					}
				}
			}
		}
	}
	return selectedPalindrome
}
