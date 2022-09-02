package arrays

// returns a slice containing the two indixes
// of the elements that when summed together
// give target as a result
func TwoSum(nums []int, target int) []int {
	var indexes []int = []int{}
	for index := 0; index < len(nums); index++ {
		if len(indexes) == 2 {
			break
		}
		for i := index + 1; i < len(nums); i++ {
			var sum = nums[index] + nums[i]
			if sum == target {
				indexes = append(indexes, index)
				indexes = append(indexes, i)
				break
			}
		}
	}
	return indexes
}
