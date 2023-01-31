package arrays

func SingleNumber(nums []int) int {
	var counter map[int]int = make(map[int]int)
	for _, value := range nums {
		_, exists := counter[value]
		if !exists {
			counter[value] = 1
		} else {
			delete(counter, value)
		}
	}
	var selected int
	for key, _ := range counter {
		selected = key
	}
	return selected
}
