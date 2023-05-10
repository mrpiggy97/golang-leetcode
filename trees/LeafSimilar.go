package trees

import (
	"reflect"
)

func GetFinalLevelsInOrder(rootNode *TreeNode) []int {
	var result [][]*TreeNode = [][]*TreeNode{[]*TreeNode{rootNode}}
	var index int = 0
	var stop bool = false
	var newResult [][]*TreeNode = [][]*TreeNode{}
	// loop though every slice until newResult and Result
	// are equal
	for !stop {
		var slice []*TreeNode = result[index]
		// loop through every node in slice
		for _, node := range slice {
			var nodeSlice []*TreeNode = []*TreeNode{}
			if node.Left == nil && node.Right == nil {
				nodeSlice = append(nodeSlice, node)
			}
			if node.Left != nil {
				nodeSlice = append(nodeSlice, node.Left)
			}
			if node.Right != nil {
				nodeSlice = append(nodeSlice, node.Right)
			}
			newResult = append(newResult, nodeSlice)
		}
		if index == len(result)-1 {
			index = 0
			if reflect.DeepEqual(result, newResult) {
				result = newResult
				break
			}
			result = newResult
			newResult = [][]*TreeNode{}
		} else {
			index += 1
		}
	}
	var finalResults []int = []int{}
	for _, slice := range result {
		finalResults = append(finalResults, slice[0].Val)
	}
	return finalResults
}

func LeafSimilar(first, second *TreeNode) bool {
	var left []int = GetFinalLevelsInOrder(first)
	var right []int = GetFinalLevelsInOrder(second)
	return reflect.DeepEqual(left, right)
}
