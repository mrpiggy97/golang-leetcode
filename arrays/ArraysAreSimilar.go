package arrays

func AreSimilar(beginningIndex int, nums1, nums2 []int) bool {
	for index := beginningIndex; index < len(nums1); index++ {
		if nums1[index] != nums2[index] {
			return false
		}
	}
	return true
}

func FindSwap(nums, nums2 []int, beginning int, value int) bool {
	var numsCopy []int = make([]int, len(nums))
	copy(numsCopy, nums)
	for i := beginning; i < len(nums); i++ {
		var currentVal int = nums[i]
		if currentVal == value {
			var previousValue int = nums[beginning]
			numsCopy[beginning] = currentVal
			numsCopy[i] = previousValue
			var slicesMatch bool = AreSimilar(beginning, numsCopy, nums2)
			if slicesMatch {
				return true
			} else {
				numsCopy[beginning] = previousValue
				numsCopy[i] = currentVal
			}
		}
	}
	return false
}

func ArraysAreSimilar(nums1, nums2 []int) bool {
	var currentIndex int = 0
	for currentIndex < len(nums1) {
		var numInNums1 int = nums1[currentIndex]
		var numInNums2 int = nums2[currentIndex]
		if numInNums1 != numInNums2 {
			var swapFound bool = FindSwap(nums1, nums2, currentIndex, numInNums2)
			if swapFound {
				return true
			} else {
				swapFound = FindSwap(nums2, nums1, currentIndex, numInNums1)
				if swapFound {
					return true
				} else {
					return false
				}
			}
		}
		currentIndex += 1
	}
	return true
}
