package linkedlists

import (
	"fmt"
	"strconv"
)

func formStr(node *ListNode) string {
	var finalStr string = ""
	var currentNode *ListNode = node
	for currentNode != nil {
		finalStr = fmt.Sprintf("%s%d", finalStr, currentNode.Val)
		currentNode = currentNode.Next
	}
	return finalStr
}

func SumNums(num1, num2 string) string {
	var leading int = 0
	var index1 int = len(num1) - 1
	var index2 int = len(num2) - 1
	var finalStr string = ""
	for index1 >= 0 && index2 >= 0 {
		var current1 string = string(num1[index1])
		var current2 string = string(num2[index2])
		digit1, _ := strconv.Atoi(current1)
		digit2, _ := strconv.Atoi(current2)
		var result int = digit1 + digit2
		if leading > 0 {
			result += 1
			leading = 0
		}
		if result >= 10 {
			result = result - 10
			leading = 1
		}
		finalStr = fmt.Sprintf("%d%s", result, finalStr)
		index1 -= 1
		index2 -= 1
	}

	if index1 >= 0 {
		for i := index1; i >= 0; i-- {
			var current string = string(num1[i])
			if leading > 0 {
				digit, _ := strconv.Atoi(current)
				digit += 1
				leading = 0
				if digit >= 10 {
					digit = digit - 10
					leading = 1
				}
				finalStr = fmt.Sprintf("%d%s", digit, finalStr)
			} else {
				finalStr = fmt.Sprintf("%s%s", current, finalStr)
			}
		}
	}

	if index2 >= 0 {
		for i := index2; i >= 0; i-- {
			var current string = string(num2[i])
			if leading > 0 {
				digit, _ := strconv.Atoi(current)
				digit += 1
				leading = 0
				if digit >= 10 {
					digit = digit - 10
					leading = 1
				}
				finalStr = fmt.Sprintf("%d%s", digit, finalStr)
			} else {
				finalStr = fmt.Sprintf("%s%s", current, finalStr)
			}
		}
	}
	if leading > 0 {
		finalStr = fmt.Sprintf("%d%s", leading, finalStr)
	}
	return finalStr
}

func SumLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var str1 string = formStr(l1)
	var str2 string = formStr(l2)
	var numStr string = SumNums(str1, str2)
	var head *ListNode = nil
	var currentNode *ListNode = nil
	for index := range numStr {
		var current string = string(numStr[index])
		digit, _ := strconv.Atoi(current)
		if head == nil {
			head = &ListNode{
				Val:  digit,
				Next: nil,
			}
			currentNode = head
		} else {
			var newNode *ListNode = &ListNode{
				Val:  digit,
				Next: nil,
			}
			currentNode.Next = newNode
			currentNode = currentNode.Next
		}
	}
	return head
}
