package arrays

func Intersect(nums1, nums2 []int) []int {
	var nums1Count map[int]int = make(map[int]int)
	var nums2Count map[int]int = make(map[int]int)
	var intersection []int = []int{}
	for _, val := range nums1 {
		count, exists := nums1Count[val]
		if !exists {
			nums1Count[val] = 1
		} else {
			nums1Count[val] = count + 1
		}
	}

	for _, val := range nums2 {
		count, exists := nums2Count[val]
		if !exists {
			nums2Count[val] = 1
		} else {
			nums2Count[val] = count + 1
		}
	}

	for val, count := range nums1Count {
		countInNums2, exists := nums2Count[val]
		if exists {
			var finalCount int = -1
			if count > countInNums2 {
				finalCount = countInNums2
			}
			if countInNums2 > count {
				finalCount = count
			}
			if count == countInNums2 {
				finalCount = count
			}
			for i := 0; i < finalCount; i++ {
				intersection = append(intersection, val)
			}
		}
	}
	return intersection
}
