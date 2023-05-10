package arrays

func getTommorrowsCell(cell []int) []int {
	var tommorowsCell []int = []int{}
	for index := range cell {
		if index == 0 {
			tommorowsCell = append(tommorowsCell, 0)
		}
		if index == len(cell)-1 {
			tommorowsCell = append(tommorowsCell, 0)
		}
		if index > 0 && index < len(cell)-1 {
			var left int = cell[index-1]
			var right int = cell[index+1]
			if left == right {
				tommorowsCell = append(tommorowsCell, 1)
			} else {
				tommorowsCell = append(tommorowsCell, 0)
			}
		}
	}
	return tommorowsCell
}

func PrisonAfterNDays(cells []int, days int) []int {
	var currentCell []int = make([]int, len(cells))
	copy(currentCell, cells)
	for i := 0; i <= days; i++ {
		currentCell = getTommorrowsCell(currentCell)
	}
	return currentCell
}
