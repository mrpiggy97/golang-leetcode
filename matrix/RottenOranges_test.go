package matrix_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/matrix"
)

func testRottenOranges(testCase *testing.T) {
	var firstGrid [][]int = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	var secondGrid [][]int = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	var thirdGrid [][]int = [][]int{{1}, {0}}
	var expectedResult int = -1
	var result int = matrix.OrangesRotting(firstGrid)
	if expectedResult != result {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
	expectedResult = 4
	result = matrix.OrangesRotting(secondGrid)
	if expectedResult != result {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
	expectedResult = -1
	result = matrix.OrangesRotting(thirdGrid)
	if expectedResult != result {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestRottenOranges(testCase *testing.T) {
	testCase.Run("action=test-rotten-oranges", testRottenOranges)
}
