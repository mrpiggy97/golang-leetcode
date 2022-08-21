package trees_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/trees"
)

func testBasicFunctionality(testCase *testing.T) {
	var firstTree *trees.TreeNode = trees.NewTreeNode()
	var firstNode int = 33
	var secondNode int = 40
	var thirdNode int = 20
	var fourthNode int = 25
	var fifthNode int = 35

	firstTree.Insert(firstNode)
	firstTree.Insert(secondNode)
	firstTree.Insert(thirdNode)
	firstTree.Insert(fourthNode)
	firstTree.Insert(fifthNode)

	if firstTree.Root.Value != firstNode {
		testCase.Errorf("expected %d, got %d instead", firstNode, firstTree.Root.Value)
	}

	if firstTree.Root.Left.Value != thirdNode {
		testCase.Errorf("expected %d, got %d instead", thirdNode, firstTree.Root.Left.Value)
	}

	if firstTree.Root.Right.Value != secondNode {
		testCase.Errorf("expected %d, got %d instead", secondNode, firstTree.Root.Right.Value)
	}

	if firstTree.Root.Left.Right.Value != fourthNode {
		testCase.Errorf("expected %d, got %d instead", fourthNode, firstTree.Root.Left.Right.Value)
	}

	if firstTree.Root.Right.Left.Value != fifthNode {
		testCase.Errorf("expected %d, got %d instead", fifthNode, firstTree.Root.Right.Left.Value)
	}
}

func TestMyTree(testCase *testing.T) {
	testCase.Run("action=tree-basic-functionality", testBasicFunctionality)
}
