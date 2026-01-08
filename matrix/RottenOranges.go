package matrix

import "fmt"

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
				if i-1 >= 0 {
					if grid[outerIndex][i-1] == 1 {
						grid[outerIndex][i-1] = 2
						orangesTurned += 1
					}
				}
				if i+1 < len(currentSlice) {
					if grid[outerIndex][i+1] == 1 {
						grid[outerIndex][i+1] = 2
						orangesTurned += 1
					}
				}
				// turn oranges up and below rotten orange
				if outerIndex-1 >= 0 {
					if grid[outerIndex-1][i] == 1 {
						grid[outerIndex-1][i] = 2
						orangesTurned += 1
					}
				}
				if outerIndex+1 < len(grid) {
					if outerIndex == 1 {
						fmt.Println(grid[outerIndex+1][i])
					}
					if grid[outerIndex+1][i] == 1 {
						grid[outerIndex+1][i] = 2
						orangesTurned += 1
					}
				}
			}
			// orange is 1, increase unturned count by 1
			if currentSlice[i] == 1 {
				orangesFound += 1
			}
		}
		minutes += 1
		fmt.Println(grid)
		outerIndex += 1
		if outerIndex == len(grid)-1 {
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
