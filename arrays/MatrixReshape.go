package arrays

func getElementsInMatrix(matrix [][]int) []int {
	var elements []int = []int{}
	for _, slice := range matrix {
		elements = append(elements, slice...)
	}
	return elements
}

func FillNewMatrix(elements []int, numberOfRows int, numberOfColumns int) [][]int {
	var newMatrix [][]int = make([][]int, numberOfRows)
	var elementIndex int = 0
	for index := range newMatrix {
		var newSlice []int = make([]int, numberOfColumns)
		for i := 0; i < len(newSlice); i++ {
			newSlice[i] = elements[elementIndex]
			elementIndex = elementIndex + 1
		}
		newMatrix[index] = newSlice
	}
	return newMatrix
}

func MatrixReshape(matrix [][]int, rows int, columns int) [][]int {
	var matrixElements []int = getElementsInMatrix(matrix)
	if len(matrixElements)%rows != 0 || len(matrixElements)/rows != columns {
		return matrix
	}
	var newMatrix [][]int = FillNewMatrix(matrixElements, rows, columns)
	return newMatrix
}
