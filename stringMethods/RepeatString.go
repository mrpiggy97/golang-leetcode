package stringMethods

func RepeatString(repetitions int, stringInstance string) string {
	var newString string = ""
	for i := 0; i < repetitions; i++ {
		newString = newString + stringInstance
	}
	return newString
}
