package arithmetic

import (
	"fmt"
	"strconv"
)

func Reverse(num int) int {
	var numStr string = fmt.Sprintf("%d", num)
	var newStr string = ""
	for i := len(numStr) - 1; i >= 0; i-- {
		var digitStr string = string(numStr[i])
		if i == 0 {
			if digitStr == "-" {
				newStr = fmt.Sprintf("%s%s", digitStr, newStr)
			} else {
				newStr = fmt.Sprintf("%s%s", newStr, digitStr)
			}
		}
		if i > 0 {
			newStr = fmt.Sprintf("%s%s", newStr, digitStr)
		}
	}
	newNum, err := strconv.Atoi(newStr)
	if err != nil {
		return 0
	}
	if newNum > 2147483647 || newNum < -2147483648 {
		return 0
	}
	return newNum
}
