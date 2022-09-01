package arrays

func GetHighestSumFromStartingIndex(index int, slice []int) int {
	// this is the lowest value an int can be in golang
	var highestSum int = -2147483648
	var endIndex int = len(slice)
	for endIndex != index {
		var sum int = 0
		for i := index; i < endIndex; i++ {
			sum = sum + slice[i]
			if sum < 0 {
				break
			}
		}
		if sum > highestSum {
			highestSum = sum
		}

		endIndex = endIndex - 1
	}
	return highestSum
}

// get the slice with highest sum
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var highestSum int = -2147483648
	for index := range nums {
		var currentSum int = GetHighestSumFromStartingIndex(index, nums)
		if currentSum > highestSum {
			highestSum = currentSum
		}
	}
	return highestSum
}
