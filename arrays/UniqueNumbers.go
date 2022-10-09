package arrays

func UniqueNumbers(slice []int) []int {
	var numbers map[int]int = make(map[int]int)
	var uniqueNumbers []int = []int{}
	for _, num := range slice {
		count, exists := numbers[num]
		if !exists {
			numbers[num] = 1
		} else {
			numbers[num] = count + 1
		}
	}
	for num, count := range numbers {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}
	return uniqueNumbers
}
