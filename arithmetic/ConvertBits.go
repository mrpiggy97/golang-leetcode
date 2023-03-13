package arithmetic

import (
	"fmt"
	"strconv"
)

// Convert to binary and return that binary reversed
func ConvertAndReverse(bits uint32) uint32 {
	var num uint32 = bits
	var resultStr string = ""
	for num > 0 {
		var remainder uint32 = num % 2
		resultStr = resultStr + fmt.Sprintf("%d", remainder)
		num = num / 2
	}
	if len(resultStr) < 32 {
		for len(resultStr) < 32 {
			resultStr = resultStr + "0"
		}
	}
	fmt.Println(resultStr)
	convertedResult, _ := strconv.ParseInt(resultStr, 2, 64)
	return uint32(convertedResult)
}
