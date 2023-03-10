package arithmetic

import (
	"fmt"
	"strconv"
	"strings"
)

func MyAtoi(str string) int {
	var strWithoutSpace string = strings.TrimSpace(str)
	var cleanStr string = ""
	for index := range strWithoutSpace {
		var currentByte byte = strWithoutSpace[index]
		if currentByte > 47 && currentByte < 58 {
			cleanStr = fmt.Sprintf("%s%s", cleanStr, string(currentByte))
		} else {
			if (currentByte == 45 || currentByte == 43) && index == 0 {
				if currentByte == 45 {
					cleanStr = fmt.Sprintf("%s%s", string(currentByte), cleanStr)
				}
			} else {
				break
			}
		}
	}
	newNum, _ := strconv.Atoi(cleanStr)
	if newNum < -2147483648 {
		return -2147483648
	}
	if newNum > 2147483647 {
		return 2147483647
	}
	return newNum
}
