package matrix

func OrangesRotting(grid [][]int) int {
	var minutes = 0
	var outerIndex int = 0
	var orangesTurned int = 0
	var orangesFound int = 0
	for outerIndex < len(grid) {
		var currentSlice []int = grid[outerIndex]
		// loop through currentSlice
		for i := 0; i < len(currentSlice); i++ {
			// turn oranges
			if currentSlice[i] == 2 {
				// turn oranges next to current rotten orange
				var freshOrangeTurned bool = false
				if i-1 >= 0 {
					if grid[outerIndex][i-1] == 1 {
						grid[outerIndex][i-1] = 2
						orangesTurned += 1
						freshOrangeTurned = true
					}
				}
				if i+1 < len(currentSlice) {
					if grid[outerIndex][i+1] == 1 {
						grid[outerIndex][i+1] = 2
						orangesTurned += 1
						freshOrangeTurned = true
					}
				}
				// turn oranges up and below rotten orange
				if outerIndex-1 >= 0 {
					if grid[outerIndex-1][i] == 1 {
						grid[outerIndex-1][i] = 2
						orangesTurned += 1
						freshOrangeTurned = true
					}
				}
				if outerIndex+1 < len(grid) {
					if grid[outerIndex+1][i] == 1 {
						grid[outerIndex+1][i] = 2
						orangesTurned += 1
						freshOrangeTurned = true
					}
				}
				if freshOrangeTurned {
					minutes += 1
				}
			}
			// orange is 1, increase unturned count by 1
			if currentSlice[i] == 1 {
				orangesFound = orangesFound + 1
			}
		}
		outerIndex += 1
		// we traversed all slices in grid
		if outerIndex == len(grid) {
			// ends loop
			if orangesFound > 0 && orangesTurned == 0 {
				return -1
			}
			// ends loop by letting loop reach condition
			if orangesFound == 0 && orangesTurned == 0 {
				outerIndex += 1
			}
			// retarts loop
			if orangesFound > 0 && orangesTurned > 0 {
				outerIndex = 0
				orangesFound = 0
				orangesTurned = 0
			}
		}
	}
	return minutes
}
