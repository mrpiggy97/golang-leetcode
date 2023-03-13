package arrays

import "sort"

func MissingNumber(nums []int) int {
	var previousNum int = 0
	sort.Ints(nums)
	for index := range nums {
		var currentNum int = nums[index]
		if index == 0 {
			if currentNum != 0 {
				return 0
			}
			previousNum = currentNum
		}
		if index > 0 {
			var difference int = currentNum - previousNum
			if difference > 1 {
				return currentNum - 1
			}
			previousNum = currentNum
		}
	}
	return len(nums)
}
