package arrays

func MaxArrea(height []int) int {
	var leftIndex int = 0
	var rightIndex int = len(height) - 1
	var maxArea int = 0
	for leftIndex < rightIndex {
		var currentLeft int = height[leftIndex]
		var currentRight int = height[rightIndex]
		var finalHeight int
		if currentLeft > currentRight {
			finalHeight = currentRight
		}
		if currentRight > currentLeft {
			finalHeight = currentLeft
		}
		if currentLeft == currentRight {
			finalHeight = currentLeft
		}
		var currentArea int = finalHeight * (rightIndex - leftIndex)
		if currentArea > maxArea {
			maxArea = currentArea
		}
		if currentLeft == currentRight {
			leftIndex += 1
			rightIndex -= 1
		}
		if currentLeft < currentRight {
			leftIndex += 1
		}
		if currentRight < currentLeft {
			rightIndex -= 1
		}
	}
	return maxArea
}
