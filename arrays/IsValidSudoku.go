package arrays

func SliceIsValid(slice []byte) bool {
	var values map[byte]bool = make(map[byte]bool)
	for _, value := range slice {
		_, exists := values[value]
		if exists {
			return false
		} else {
			values[value] = false
		}
	}
	return true
}

func GetSliceFromIndex(sudoku [][]byte, index int) []byte {
	var elementsWithSameIndex []byte = []byte{}
	for _, slice := range sudoku {
		elementsWithSameIndex = append(elementsWithSameIndex, slice[index])
	}
	return elementsWithSameIndex
}

func GetVerticalSlices(sudoku [][]byte) [][]byte {
	var lastIndex int = len(sudoku[0])
	var verticalSlices [][]byte = make([][]byte, len(sudoku))
	for index := 0; index < lastIndex; index++ {
		var currentSlice []byte = GetSliceFromIndex(sudoku, index)
		verticalSlices[index] = currentSlice
	}
	return verticalSlices
}

func GroupSlices(sudoku [][]byte) [][][]byte {
	firstGroup := [][]byte{}
	secondGroup := [][]byte{}
	thirdGroup := [][]byte{}

	for index := 0; index < 3; index++ {
		firstGroup = append(firstGroup, sudoku[index])
	}
	for index := 3; index < 6; index++ {
		secondGroup = append(secondGroup, sudoku[index])
	}
	for index := 6; index < 9; index++ {
		thirdGroup = append(thirdGroup, sudoku[index])
	}
	return [][][]byte{firstGroup, secondGroup, thirdGroup}
}

// rearrange group into three separate slices
// that form a subboard
func GetSubBoards(group [][]byte) [][]byte {
	var firstGroup []byte = []byte{}
	var secondGroup []byte = []byte{}
	var thirdGroup []byte = []byte{}

	for _, subSlice := range group {
		for index := 0; index < 3; index++ {
			firstGroup = append(firstGroup, subSlice[index])
		}
		for index := 3; index < 6; index++ {
			secondGroup = append(secondGroup, subSlice[index])
		}
		for index := 6; index < 9; index++ {
			thirdGroup = append(thirdGroup, subSlice[index])
		}
	}
	return [][]byte{firstGroup, secondGroup, thirdGroup}
}

// check if there are repeated values in slice
// if there is a value that's repeated return true
// else return false
func CheckValues(slice []byte) bool {
	var values map[byte]bool = make(map[byte]bool)
	for _, value := range slice {
		if value != 46 {
			_, exists := values[value]
			if exists {
				return true
			} else {
				values[value] = false
			}
		}
	}
	return false
}

func CheckIfSubBoardsAreValid(subBoards [][]byte) bool {
	for _, board := range subBoards {
		var boardHasRepeatedValues bool = CheckValues(board)
		if boardHasRepeatedValues {
			return false
		}
	}
	return true
}

func IsValidSudoku(board [][]byte) bool {
	var verticalSlices [][]byte = GetVerticalSlices(board)
	var valuesRepeated bool = false
	for _, horizontalSlice := range board {
		valuesRepeated = CheckValues(horizontalSlice)
		if valuesRepeated {
			return false
		}
	}

	for _, slice := range verticalSlices {
		valuesRepeated = CheckValues(slice)
		if valuesRepeated {
			return false
		}
	}
	var groups [][][]byte = GroupSlices(board)
	for _, group := range groups {
		var subBoards [][]byte = GetSubBoards(group)
		var subBoardsAreValid bool = CheckIfSubBoardsAreValid(subBoards)
		if !subBoardsAreValid {
			return false
		}
	}
	return true
}
