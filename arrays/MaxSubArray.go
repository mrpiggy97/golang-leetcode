package arrays

func GetSliceWithHighestSum(slices [][]int) ([]int, int) {
	var sliceWithHighestValue []int = []int{}
	// this is the lowest value an int can be in golang
	var highestSum int = -2147483648
	for _, slice := range slices {
		var sum int = 0
		for _, value := range slice {
			sum = sum + value
		}
		if sum > highestSum {
			highestSum = sum
			sliceWithHighestValue = slice
		}
	}
	return sliceWithHighestValue, highestSum
}

func GetSlicesFromIndex(index int, slice []int) [][]int {
	var slices [][]int = [][]int{}
	var endIndex int = len(slice)

	for endIndex != index {
		var currentSlice []int = []int{}
		for i := index; i < endIndex; i++ {
			currentSlice = append(currentSlice, slice[i])
		}
		slices = append(slices, currentSlice)
		endIndex = endIndex - 1
	}

	return slices
}

// get the slice with highest sum
func MaxSubArray(nums []int) int {
	var slicesWithHighestSum [][]int = [][]int{}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for index := range nums {
		var currentSlice [][]int = GetSlicesFromIndex(index, nums)
		highestValueSlice, _ := GetSliceWithHighestSum(currentSlice)
		slicesWithHighestSum = append(slicesWithHighestSum, highestValueSlice)
	}

	_, sum := GetSliceWithHighestSum(slicesWithHighestSum)
	return sum
}
