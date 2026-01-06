package arrays

func MaxNumOfMarkedIndices(indices []int) int {
	var count int = 0
	var jIndex int = len(indices) - 1
	var iIndex int = len(indices) / 2

	// first we need to sort the array from smallest to biggest
	var sortedSlice []int = SortSlice(indices)
}
