package trees

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Right: nil,
		Left:  nil,
	}
}

type Tree struct {
	Root *TreeNode
}

func (tree *Tree) getIndexOfValue(value *TreeNode, slice []*TreeNode) int {
	var correctIndex int = -1
	for index, val := range slice {
		if val == value {
			correctIndex = index
			break
		}
	}
	return correctIndex
}

func (tree *Tree) insertAt(index int, newValue *TreeNode, slice []*TreeNode) []*TreeNode {
	var values []*TreeNode = []*TreeNode{}
	if index == 0 {
		values = append(values, newValue)
		values = append(values, slice...)
	} else {
		var valuesUpToIndex []*TreeNode = []*TreeNode{}
		var valuesFromIndexOnward []*TreeNode = []*TreeNode{}
		for i := 0; i < index; i++ {
			valuesUpToIndex = append(valuesUpToIndex, slice[i])
		}
		for a := index; a < len(slice); a++ {
			valuesFromIndexOnward = append(valuesFromIndexOnward, slice[a])
		}
		valuesUpToIndex = append(valuesUpToIndex, newValue)
		values = append(values, valuesUpToIndex...)
		values = append(values, valuesFromIndexOnward...)
	}
	return values
}

func (tree *Tree) Insert(value int) {

	if tree.Root != nil {
		var currentNode *TreeNode = tree.Root
		var stop bool = false
		for !stop {
			if value == currentNode.Val {
				fmt.Println("a node with that value already exists")
			}
			if value > currentNode.Val {
				if currentNode.Right == nil {
					currentNode.Right = NewNode(value)
					stop = true
				}
				if currentNode.Right != nil {
					currentNode = currentNode.Right
				}
			}

			if value < currentNode.Val {
				if currentNode.Left == nil {
					currentNode.Left = NewNode(value)
					stop = true
				}
				if currentNode.Left != nil {
					currentNode = currentNode.Left
				}
			}
		}
	}

	if tree.Root == nil {
		tree.Root = NewNode(value)
	}
}

func (tree *Tree) Search(value int) *TreeNode {
	var currentNode *TreeNode = tree.Root
	var stop bool = false
	for !stop {
		if currentNode.Val == value {
			stop = true
		}

		if value > currentNode.Val {
			if currentNode.Right == nil {
				fmt.Println("no node with that value exists")
				return nil
			}
			if currentNode.Right != nil {
				currentNode = currentNode.Right
			}
		}

		if value < currentNode.Val {
			if currentNode.Left == nil {
				fmt.Println("no node with that value exists")
				return nil
			}
			if currentNode.Left != nil {
				currentNode = currentNode.Left
			}
		}
	}
	return currentNode
}

func (tree *Tree) getNodesInOrder(node *TreeNode) []int {
	var nodesToVisit []*TreeNode = []*TreeNode{node}
	var nodesInOrder []*TreeNode = []*TreeNode{node}

	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *TreeNode = nodesToVisit[i]
		var currentIndex int = tree.getIndexOfValue(currentNode, nodesInOrder)
		var leftNode *TreeNode = currentNode.Left
		var rightNode *TreeNode = currentNode.Right

		if leftNode != nil {
			nodesToVisit = append(nodesToVisit, leftNode)
			nodesInOrder = tree.insertAt(currentIndex, leftNode, nodesInOrder)
			currentIndex = currentIndex + 1
		}
		if rightNode != nil {
			nodesToVisit = append(nodesToVisit, rightNode)
			if currentIndex == len(nodesInOrder)-1 {
				nodesInOrder = append(nodesInOrder, rightNode)
			} else {
				currentIndex = currentIndex + 1
				nodesInOrder = tree.insertAt(currentIndex, rightNode, nodesInOrder)
			}
		}
	}

	var valuesOfNodesInOrder []int = []int{}
	for _, pointer := range nodesInOrder {
		valuesOfNodesInOrder = append(valuesOfNodesInOrder, pointer.Val)
	}
	return valuesOfNodesInOrder
}

func (tree *Tree) getNodesInPreOrder(node *TreeNode) []int {
	var nodesInPreOrder []*TreeNode = []*TreeNode{node}
	var nodesToVisit []*TreeNode = []*TreeNode{node}
	for index := 0; index < len(nodesToVisit); index++ {
		var currentNode *TreeNode = nodesToVisit[index]
		var indexOfCurrentNode int = tree.getIndexOfValue(currentNode, nodesInPreOrder)
		var leftNode *TreeNode = currentNode.Left
		var rightNode *TreeNode = currentNode.Right
		var isEndIndex bool = indexOfCurrentNode == len(nodesInPreOrder)-1
		if leftNode != nil {
			if isEndIndex {
				nodesInPreOrder = append(nodesInPreOrder, leftNode)
				nodesToVisit = append(nodesToVisit, leftNode)
			} else {
				indexOfCurrentNode = indexOfCurrentNode + 1
				nodesInPreOrder = tree.insertAt(indexOfCurrentNode, leftNode, nodesInPreOrder)
				nodesToVisit = append(nodesToVisit, leftNode)
			}
		}
		indexOfCurrentNode = indexOfCurrentNode + 1
		if rightNode != nil {
			if isEndIndex {
				nodesInPreOrder = append(nodesInPreOrder, rightNode)
				nodesToVisit = append(nodesToVisit, rightNode)
			} else {
				nodesInPreOrder = tree.insertAt(indexOfCurrentNode, rightNode, nodesInPreOrder)
				nodesToVisit = append(nodesToVisit, rightNode)
			}
		}
	}
	var nodes []int = make([]int, len(nodesInPreOrder))
	for index, node := range nodesInPreOrder {
		nodes[index] = node.Val
	}
	return nodes
}

func (tree *Tree) getNodesInPostOrder(node *TreeNode) []int {
	var nodesInPostOrder []*TreeNode = []*TreeNode{node}
	var nodesToVisit []*TreeNode = []*TreeNode{node}

	for index := 0; index < len(nodesToVisit); index++ {
		var currentNode *TreeNode = nodesToVisit[index]
		var leftNode *TreeNode = currentNode.Left
		var rightNode *TreeNode = currentNode.Right
		var currentNodeIndex int = tree.getIndexOfValue(currentNode, nodesInPostOrder)

		if leftNode != nil {
			nodesInPostOrder = tree.insertAt(currentNodeIndex, leftNode, nodesInPostOrder)
			nodesToVisit = append(nodesToVisit, leftNode)
			currentNodeIndex = currentNodeIndex + 1
		}

		if rightNode != nil {
			nodesInPostOrder = tree.insertAt(currentNodeIndex, rightNode, nodesInPostOrder)
			nodesToVisit = append(nodesToVisit, rightNode)
		}
	}
	var nodes []int = make([]int, len(nodesInPostOrder))
	for index, node := range nodesInPostOrder {
		nodes[index] = node.Val
	}
	return nodes
}

func (tree *Tree) InOrderTraverse() []int {
	if tree.Root == nil {
		fmt.Println("tree is empty")
		return []int{}
	}

	var nodes []int = []int{}

	if tree.Root.Left != nil {
		nodes = tree.getNodesInOrder(tree.Root.Left)
		nodes = append(nodes, tree.Root.Val)
	} else {
		nodes = append(nodes, tree.Root.Val)
	}

	if tree.Root.Right != nil {
		var rightNodes []int = tree.getNodesInOrder(tree.Root.Right)
		nodes = append(nodes, rightNodes...)
	}
	return nodes
}

func (tree *Tree) PreOrderTraverse() []int {
	if tree.Root == nil {
		return []int{}
	}
	var nodes []int = []int{tree.Root.Val}
	var leftNode *TreeNode = tree.Root.Left
	var rightNode *TreeNode = tree.Root.Right

	if leftNode != nil {
		nodes = append(nodes, tree.getNodesInPreOrder(leftNode)...)
	}
	if rightNode != nil {
		nodes = append(nodes, tree.getNodesInPreOrder(rightNode)...)
	}
	return nodes
}

func (tree *Tree) PostOrderTraverse() []int {
	if tree.Root == nil {
		return []int{}
	}

	return tree.getNodesInPostOrder(tree.Root)
}

func (tree *Tree) LevelOrderTraverse() [][]int {
	if tree.Root == nil {
		return [][]int{}
	}

	var currentLevel []*TreeNode = []*TreeNode{tree.Root}
	var level []int = []int{tree.Root.Val}
	var levels [][]int = [][]int{level}
	var currentLevelLength int = len(currentLevel)

	for currentLevelLength > 0 {
		var newLevel []*TreeNode = []*TreeNode{}
		var newValues []int = []int{}
		for index := 0; index < currentLevelLength; index++ {
			var currentNode *TreeNode = currentLevel[index]
			var leftNode *TreeNode = currentNode.Left
			var rightNode *TreeNode = currentNode.Right

			if leftNode != nil {
				newLevel = append(newLevel, leftNode)
				newValues = append(newValues, leftNode.Val)
			}
			if rightNode != nil {
				newLevel = append(newLevel, rightNode)
				newValues = append(newValues, rightNode.Val)
			}
		}
		currentLevel = newLevel
		if len(newValues) > 0 {
			levels = append(levels, newValues)
		}
		currentLevelLength = len(currentLevel)
	}
	return levels
}

func NewTree() *Tree {
	return &Tree{
		Root: nil,
	}
}
