package arrays

func TwoSum2(nums []int, target int) []int {
	var indexes []int = []int{0, 0}
	for index, value := range nums {
		for i := index; i < len(nums); i++ {
			if i != index {
				var val int = nums[i]
				if value+val == target {
					indexes[0] = index
					indexes[1] = i
					break
				}
			}
		}
	}
	return indexes
}
