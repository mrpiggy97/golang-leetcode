package linkedlists

import (
	"fmt"
	"strconv"
)

// Return a string containing all the digits
// in reverse order from a given LinkedList
func GetNumberStr(node *ListNode) string {
	var current *ListNode = node
	var str string = ""
	for current != nil {
		str = fmt.Sprintf("%d%s", current.Val, str)
		current = current.Next
	}
	return str
}

// Add str1 and str2
func SumNumbers(str1, str2 string) string {
	var newStr string = ""
	var index1 int = len(str1) - 1
	var index2 int = len(str2) - 1
	var leading int = 0
	for index1 >= 0 && index2 >= 0 {
		// get based 10 digit
		var firstStr string = string(str1[index1])
		var secondStr string = string(str2[index2])
		// convert based 10 digit
		num1, _ := strconv.Atoi(firstStr)
		num2, _ := strconv.Atoi(secondStr)
		// add digits
		var result int = num1 + num2
		// if leading is bigger than zero increase result by 1
		if leading > 0 {
			result += 1
			leading = 0
		}
		// if result is bigger or equal than 10 subtract result by 10
		// and set leading to 1
		if result >= 10 {
			leading = 1
			result = result - 10
		}
		// place result at the beggining of newStr
		newStr = fmt.Sprintf("%d%s", result, newStr)
		index1 -= 1
		index2 -= 1
	}

	// if either str1 or str2 are bigger than eachother
	// place at the beggining all the remaining string characters of whoever
	// has the biggest length between str1 and str2
	// if leading continues to be 1 than add that to the current digit
	// and if the result of that is bigger or equal than 10 substract 10
	// from result and set leading back to 1
	if len(str1) > len(str2) {
		for i := index1; i >= 0; i-- {
			if leading == 1 {
				currentVal, _ := strconv.Atoi(string(str1[i]))
				currentVal += 1
				leading = 0
				if currentVal >= 10 {
					currentVal = currentVal - 10
					leading = 1
				}
				newStr = fmt.Sprintf("%d%s", currentVal, newStr)
			} else {
				newStr = fmt.Sprintf("%s%s", string(str1[i]), newStr)
			}
		}
	}

	if len(str2) > len(str1) {
		for i := index2; i >= 0; i-- {
			if leading == 1 {
				currentVal, _ := strconv.Atoi(string(str2[i]))
				currentVal += 1
				leading = 0
				if currentVal >= 10 {
					currentVal = currentVal - 10
					leading = 1
				}
				newStr = fmt.Sprintf("%d%s", currentVal, newStr)
			} else {
				newStr = fmt.Sprintf("%s%s", string(str2[i]), newStr)
			}
		}
	}
	if leading == 1 {
		newStr = fmt.Sprintf("1%s", newStr)
	}
	return newStr
}

func AddTwoNumbers(first, second *ListNode) *ListNode {
	var firstStr string = GetNumberStr(first)
	var secondStr string = GetNumberStr(second)
	var result string = SumNumbers(firstStr, secondStr)
	var newHead *ListNode = nil
	var current *ListNode = nil
	for i := len(result) - 1; i >= 0; i-- {
		var currentStr string = string(result[i])
		convertedStr, _ := strconv.Atoi(currentStr)
		if newHead == nil {
			newHead = &ListNode{
				Val:  convertedStr,
				Next: nil,
			}
			current = newHead
		} else {
			var newNode *ListNode = &ListNode{
				Val:  convertedStr,
				Next: nil,
			}
			current.Next = newNode
			current = current.Next
		}
	}
	return newHead
}
