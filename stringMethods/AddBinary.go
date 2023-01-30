package stringMethods

func SumDigits(firstElement string, secondElement string, leading string) (string, string) {
	var finalResult string = ""
	var newLeading string = ""
	if firstElement == "1" && secondElement == "1" {
		newLeading = "1"
		if leading == "1" {
			finalResult = "1"
		} else {
			finalResult = "0"
		}
	}

	if (firstElement == "0" && secondElement == "1") || (secondElement == "0" && firstElement == "1") {
		if leading == "1" {
			finalResult = "0"
			newLeading = "1"
		} else {
			finalResult = "1"
		}
	}

	if firstElement == "0" && secondElement == "0" {
		if leading == "1" {
			finalResult = "1"
		} else {
			finalResult = "0"
		}
	}
	return finalResult, newLeading
}

func AddBinary(a string, b string) string {
	var result string = ""
	var endIndex1 int = len(a) - 1
	var endIndex2 int = len(b) - 1
	var leading string = ""
	for endIndex1 >= 0 && endIndex2 >= 0 {
		var firstElement string = string(a[endIndex1])
		var secondElement string = string(b[endIndex2])
		var sum string = ""
		sum, leading = SumDigits(firstElement, secondElement, leading)
		result = sum + result
		endIndex1 = endIndex1 - 1
		endIndex2 = endIndex2 - 1
	}
	var leftOverIndex int = -1
	var lefOver string = ""
	if endIndex1 >= 0 {
		leftOverIndex = endIndex1
		lefOver = a
	}
	if endIndex2 >= 0 {
		leftOverIndex = endIndex2
		lefOver = b
	}

	for currentIndex := leftOverIndex; currentIndex >= 0; currentIndex-- {
		var currentElement string = string(lefOver[currentIndex])
		if leading == "1" {
			if currentElement == "0" {
				result = "1" + result
				leading = ""
			} else {
				result = "0" + result
			}
		} else {
			result = currentElement + result
		}
	}
	if leading == "1" {
		result = leading + result
	}
	return result
}
