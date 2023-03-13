package arrays

import (
	"sort"
)

func filterSlice(nums []int) []int {
	var newNums []int = []int{}
	for _, value := range nums {
		if value != -1 {
			newNums = append(newNums, value)
		}
	}
	sort.Ints(newNums)
	return newNums
}

func SortWithTree(nums []int) []int {
	var currentIndex int = 0
	var sortedNums []int = filterSlice(nums)
	for index, value := range nums {
		if value != -1 {
			nums[index] = sortedNums[currentIndex]
			currentIndex += 1
		}
	}
	return nums
}
