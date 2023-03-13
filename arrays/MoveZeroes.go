package arrays

func MoveZeroes(nums []int) {
	var comparison []int = []int{}
	var zeroCounter int = 0
	for _, value := range nums {
		if value == 0 {
			zeroCounter += 1
		}
		if value != 0 {
			comparison = append(comparison, value)
		}
	}

	for i := 0; i < zeroCounter; i++ {
		comparison = append(comparison, 0)
	}

	for index := range nums {
		nums[index] = comparison[index]
	}
}
