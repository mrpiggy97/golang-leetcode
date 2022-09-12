package arrays

func GetNewLevel(currentLevel []int) []int {
	var newLevel []int = make([]int, len(currentLevel)+1)
	newLevel[0] = 1
	newLevel[len(newLevel)-1] = 1
	var currentNumber int = currentLevel[0]
	var newLevelIndex int = 1
	for index := 1; index < len(currentLevel); index++ {
		var totalValue int = currentNumber + currentLevel[index]
		currentNumber = currentLevel[index]
		newLevel[newLevelIndex] = totalValue
		newLevelIndex = newLevelIndex + 1
	}
	return newLevel
}

//return the row level specified of pascal's triangle
func Generate(numRows int) [][]int {
	var pascalLevels [][]int = [][]int{{1}, {1, 1}}
	if numRows == 1 || numRows == 2 {
		return pascalLevels[0:numRows]
	}
	numRows = numRows - 2
	for numRows > 0 {
		var currentLevel []int = pascalLevels[len(pascalLevels)-1]
		var newLevel []int = GetNewLevel(currentLevel)
		pascalLevels = append(pascalLevels, newLevel)
		numRows = numRows - 1
	}
	return pascalLevels
}
