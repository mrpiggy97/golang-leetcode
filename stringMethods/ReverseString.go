package stringMethods

func ReverseString(initialString []byte) {
	var start int = 0
	var end int = len(initialString) - 1
	for start < end {
		var currentLeft byte = initialString[start]
		var currentRight byte = initialString[end]
		initialString[start] = currentRight
		initialString[end] = currentLeft
		start += 1
		end -= 1
	}
}
