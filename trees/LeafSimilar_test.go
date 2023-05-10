package trees_test

import (
	"reflect"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/trees"
)

func testGetFinalLevelInOrder(testCase *testing.T) {
	var rootNode *trees.TreeNode = trees.NewNode(3)
	var secondNode *trees.TreeNode = trees.NewNode(5)
	var thirdNode *trees.TreeNode = trees.NewNode(1)
	rootNode.Left = secondNode
	rootNode.Right = thirdNode

	var fourthNode *trees.TreeNode = trees.NewNode(6)
	var fifthNode *trees.TreeNode = trees.NewNode(2)
	secondNode.Left = fourthNode
	secondNode.Right = fifthNode

	var sixthNode *trees.TreeNode = trees.NewNode(9)
	var seventhNode *trees.TreeNode = trees.NewNode(8)
	thirdNode.Left = sixthNode
	thirdNode.Right = seventhNode

	var eigthNode *trees.TreeNode = trees.NewNode(7)
	var ninthNode *trees.TreeNode = trees.NewNode(4)
	fifthNode.Left = eigthNode
	fifthNode.Right = ninthNode
	var expectedResult []int = []int{6, 7, 4, 9, 8}
	var result []int = trees.GetFinalLevelsInOrder(rootNode)
	if !reflect.DeepEqual(result, expectedResult) {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testLeafSimilar(testCase *testing.T) {
	var firstNode *trees.TreeNode = trees.NewNode(5)
	var secondNode *trees.TreeNode = trees.NewNode(4)
	var thirdNode *trees.TreeNode = trees.NewNode(6)
	firstNode.Left = secondNode
	firstNode.Right = thirdNode

	var fourthNode *trees.TreeNode = trees.NewNode(7)
	var fifthNode *trees.TreeNode = trees.NewNode(4)
	var sixthNode *trees.TreeNode = trees.NewNode(6)
	fourthNode.Left = fifthNode
	fourthNode.Right = sixthNode

	var expectedResult bool = true
	var result bool = trees.LeafSimilar(firstNode, fourthNode)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestLeafSimilar(testCase *testing.T) {
	testCase.Run("action=test-get-final-level-in-order", testGetFinalLevelInOrder)
	testCase.Run("action=test-leaf-similar", testLeafSimilar)
}
