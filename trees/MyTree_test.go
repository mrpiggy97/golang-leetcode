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

func testPostOrderTraverse(testCase *testing.T) {
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
	var result []int = myTree.PostOrderTraverse()
	var expectedResult []int = []int{4, 2, 9, 8, 7, 12, 11, 19, 29, 33, 50, 40, 21, 20}
	for index, node := range result {
		if node != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", node, expectedResult[index])
		}
	}
}

func testLevelOrderTraverse(testCase *testing.T) {
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
	var result [][]int = myTree.LevelOrderTraverse()
	fmt.Println(result)
	var expectedResult [][]int = [][]int{{20}, {19, 21}, {11, 40}, {7, 12, 33, 50}, {2, 8, 29}, {4, 9}}
	if len(result) != len(expectedResult) {
		testCase.Errorf("expected length of result to be %d, got %d", len(expectedResult), len(result))
	}
}

func testZigZagLevelTraverse(testCase *testing.T) {
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
	var result [][]int = myTree.ZigZagLevelTraverse()
	var expectedResult [][]int = [][]int{{20}, {21, 19}, {11, 40}, {50, 33, 12, 7}, {2, 8, 29}, {9, 4}}
	for index, slice := range result {
		var sliceComparing []int = expectedResult[index]
		for nodeIndex, node := range slice {
			if node != sliceComparing[nodeIndex] {
				testCase.Errorf("expected %d to be %d", node, sliceComparing[nodeIndex])
			}
		}
	}
}

func testGetDepth(testCase *testing.T) {
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
	var result int = myTree.GetDepth()
	var expectedResult int = 6
	if result != expectedResult {
		testCase.Errorf("expected depth of tree to be %d, got %d instead", expectedResult, result)
	}
}

func testLevelIsMirror(testCase *testing.T) {
	var myTree *trees.Tree = trees.NewTree()
	var firstNode *trees.TreeNode = &trees.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	var secondNode *trees.TreeNode = &trees.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	var level []*trees.TreeNode = []*trees.TreeNode{nil, firstNode, secondNode, nil}
	var result bool = myTree.LevelIsMirror(level)
	var expectedResult bool = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	level = []*trees.TreeNode{nil, firstNode, nil, secondNode}
	result = myTree.LevelIsMirror(level)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	firstNode.Val = 4
	secondNode.Val = 3
	var thirdNode *trees.TreeNode = &trees.TreeNode{
		Val:   3,
		Right: nil,
		Left:  nil,
	}
	var fourthNode *trees.TreeNode = &trees.TreeNode{
		Val:   4,
		Right: nil,
		Left:  nil,
	}
	level = []*trees.TreeNode{firstNode, secondNode, thirdNode, fourthNode}
	result = myTree.LevelIsMirror(level)
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testIsMirror(testCase *testing.T) {
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
	var result bool = myTree.IsMirror()
	var expectedResult bool = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testHasPathSum(testCase *testing.T) {
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
	var result bool = myTree.HasPathSum(114)
	var expectedResult bool = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	result = myTree.HasPathSum(1000)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}
func TestMyTree(testCase *testing.T) {
	testCase.Run("action=tree-basic-functionality", testBasicFunctionality)
	testCase.Run("action=tree-search", testSearch)
	testCase.Run("action=test-in-order-traverse", testInOrderTraverse)
	testCase.Run("action=test-pre-order-traverse", testPreOrderTraverse)
	testCase.Run("action=test-post-order-traverse", testPostOrderTraverse)
	testCase.Run("action=test-level-order-traverse", testLevelOrderTraverse)
	testCase.Run("action=test-zigzag-level-traverse", testZigZagLevelTraverse)
	testCase.Run("action=test-get-depth", testGetDepth)
	testCase.Run("action=test-level-is-mirror", testLevelIsMirror)
	testCase.Run("action=test-is-mirror", testIsMirror)
	testCase.Run("action=test-has-path-sum", testHasPathSum)
}
