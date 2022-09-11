package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testFillNewMatrix(testCase *testing.T) {
	var testElements []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var numberOfRows int = 2
	var numberOfColumns int = 6
	var newMatrix [][]int = arrays.FillNewMatrix(testElements, numberOfRows, numberOfColumns)
	fmt.Println(newMatrix)
	if len(newMatrix) != 2 {
		testCase.Errorf("expected length of new matrix to be %d, got %d instead", numberOfRows, len(newMatrix))
	}
}

func testMatrixReshape(testCase *testing.T) {
	var testingMatrix [][]int = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}}
	var newMatrix [][]int = arrays.MatrixReshape(testingMatrix, 2, 6)
	if len(newMatrix) != 2 {
		testCase.Errorf("expected length to be %d, got %d instead", 2, len(newMatrix))
	}
	if len(newMatrix[0]) != 6 {
		testCase.Errorf("expected length of row to be %d, got %d instead", 6, len(newMatrix[0]))
	}
	testingMatrix = [][]int{{1, 2}, {3, 4}}
	newMatrix = arrays.MatrixReshape(testingMatrix, 2, 4)
	if len(testingMatrix) != 2 {
		testCase.Errorf("expected length of row to be %d, got %d instead", 6, len(newMatrix[0]))
	}
}

func TestMatrixReshape(testCase *testing.T) {
	testCase.Run("action=test-fill-new-matrix", testFillNewMatrix)
	testCase.Run("action=test-matrix-reshape", testMatrixReshape)
}
