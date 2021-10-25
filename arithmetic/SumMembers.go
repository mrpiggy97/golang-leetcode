package arithmetic

import (
	"fmt"
	"strconv"
	"strings"
)

func SumMembers(number int) int {
	var stringNumber = fmt.Sprintf("%v", number)
	var separatedString []string = strings.Split(stringNumber, "")
	var copyNumber int = 0
	for len(separatedString) >= 2 {
		copyNumber = 0
		for _, digit := range separatedString {
			convertedDigit, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}

			copyNumber += convertedDigit
		}
		stringNumber = fmt.Sprintf("%v", copyNumber)
		separatedString = strings.Split(stringNumber, "")
	}

	return copyNumber
}
