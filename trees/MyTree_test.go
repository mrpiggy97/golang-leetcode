package trees_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/trees"
)

func testBasicFunctionality(testCase *testing.T) {
	var firstTree *trees.Tree = trees.NewTree()
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

	if firstTree.Root.Val != firstNode {
		testCase.Errorf("expected %d, got %d instead", firstNode, firstTree.Root.Val)
	}

	if firstTree.Root.Left.Val != thirdNode {
		testCase.Errorf("expected %d, got %d instead", thirdNode, firstTree.Root.Left.Val)
	}

	if firstTree.Root.Right.Val != secondNode {
		testCase.Errorf("expected %d, got %d instead", secondNode, firstTree.Root.Right.Val)
	}

	if firstTree.Root.Left.Right.Val != fourthNode {
		testCase.Errorf("expected %d, got %d instead", fourthNode, firstTree.Root.Left.Right.Val)
	}

	if firstTree.Root.Right.Left.Val != fifthNode {
		testCase.Errorf("expected %d, got %d instead", fifthNode, firstTree.Root.Right.Left.Val)
	}
}

func testSearch(testCase *testing.T) {
	var firstTree *trees.Tree = trees.NewTree()
	var firstNode int = 33
	var secondNode int = 40
	var thirdNode int = 20
	var fourthNode int = 25
	var fifthNode int = 35

	var nodesToSearch []int = []int{firstNode, secondNode, thirdNode, fourthNode, fifthNode}

	firstTree.Insert(firstNode)
	firstTree.Insert(secondNode)
	firstTree.Insert(thirdNode)
	firstTree.Insert(fourthNode)
	firstTree.Insert(fifthNode)

	for _, node := range nodesToSearch {
		var result *trees.TreeNode = firstTree.Search(node)
		if result.Val != node {
			testCase.Errorf("expected %d, got %d instead", node, result.Val)
		}
	}
}

func testInOrderTraverse(testCase *testing.T) {
	var myTree *trees.Tree = trees.NewTree()
	myTree.Insert(20)
	myTree.Insert(19)
	myTree.Insert(21)
	myTree.Insert(11)
	myTree.Insert(40)
	myTree.Insert(7)
	myTree.Insert(12)
	myTree.Insert(33)
	myTree.Insert(50)
	myTree.Insert(29)
	myTree.Insert(2)
	myTree.Insert(8)
	myTree.Insert(9)
	myTree.Insert(4)
	var nodesInOrder []int = myTree.InOrderTraverse()
	fmt.Println(nodesInOrder)
	var expectedOrder []int = []int{2, 4, 7, 8, 9, 11, 12, 19, 20, 21, 29, 33, 40, 50}

	for index := range nodesInOrder {
		var node int = nodesInOrder[index]
		var expectedNode int = expectedOrder[index]
		if node != expectedNode {
			testCase.Errorf("node %d is not equal to %d", node, expectedNode)
		}
	}
}

func testPreOrderTraverse(testCase *testing.T) {
	var myTree *trees.Tree = trees.NewTree()
	myTree.Insert(20)
	myTree.Insert(19)
	myTree.Insert(21)
	myTree.Insert(11)
	myTree.Insert(40)
	myTree.Insert(7)
	myTree.Insert(12)
	myTree.Insert(33)
	myTree.Insert(50)
	myTree.Insert(29)
	myTree.Insert(2)
	myTree.Insert(8)
	myTree.Insert(9)
	myTree.Insert(4)
	var nodesInPreOrder []int = myTree.PreOrderTraverse()
	var expectedResult []int = []int{20, 19, 11, 7, 2, 4, 8, 9, 12, 21, 40, 33, 29, 50}
	fmt.Println("nodes in pre order ", nodesInPreOrder)
	for index, node := range nodesInPreOrder {
		if node != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", node, expectedResult[index])
		}
	}
}

func TestMyTree(testCase *testing.T) {
	testCase.Run("action=tree-basic-functionality", testBasicFunctionality)
	testCase.Run("action=tree-search", testSearch)
	testCase.Run("action=test-in-order-traverse", testInOrderTraverse)
	testCase.Run("action=test-pre-order-traverse", testPreOrderTraverse)
}
