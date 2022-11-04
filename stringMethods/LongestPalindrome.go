package stringMethods

func IsPalindrome(str string) bool {
	var tailIndex int = len(str) - 1
	var splitIndex int = int(len(str) / 2)
	for index := 0; index < splitIndex; index++ {
		if str[index] != str[tailIndex] {
			return false
		}
		tailIndex = tailIndex - 1
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
				if length <= len(selectedPalindrome) {
					break
				}
				var newStr string = str[index:tail] + string(str[tail])
				var isPalindrome bool = IsPalindrome(newStr)
				if isPalindrome {
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
