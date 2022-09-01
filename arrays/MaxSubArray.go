package arrays

func GetHighestSumFromStartingIndex(slice []int) int {
	// this is the lowest value an int can be in golang
	var highestSum int = -2147483648
	for index := 0; index < len(slice); index++ {
		var sum int = 0
		for i := index; i < len(slice); i++ {
			var currentNumber int = slice[i]
			sum = sum + currentNumber
			if sum > highestSum {
				highestSum = sum
			}

			if currentNumber < 0 && sum < 0 {
				break
			}
		}
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
	var highestSum int = GetHighestSumFromStartingIndex(nums)
	return highestSum
}
