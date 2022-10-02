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
	var nodesInPostOrder []*TreeNode = []*TreeNode{node}
	var nodesToVisit []*TreeNode = []*TreeNode{node}
	for index := 0; index < len(nodesToVisit); index++ {
		var currentNode *TreeNode = nodesToVisit[index]
		var indexOfCurrentNode int = tree.getIndexOfValue(currentNode, nodesInPostOrder)
		var leftNode *TreeNode = currentNode.Left
		var rightNode *TreeNode = currentNode.Right
		var isEndIndex bool = indexOfCurrentNode == len(nodesInPostOrder)-1
		if leftNode != nil {
			if isEndIndex {
				nodesInPostOrder = append(nodesInPostOrder, leftNode)
				nodesToVisit = append(nodesToVisit, leftNode)
			} else {
				indexOfCurrentNode = indexOfCurrentNode + 1
				nodesInPostOrder = tree.insertAt(indexOfCurrentNode, leftNode, nodesInPostOrder)
				nodesToVisit = append(nodesToVisit, leftNode)
			}
		}
		indexOfCurrentNode = indexOfCurrentNode + 1
		if rightNode != nil {
			if isEndIndex {
				nodesInPostOrder = append(nodesInPostOrder, rightNode)
				nodesToVisit = append(nodesToVisit, rightNode)
			} else {
				nodesInPostOrder = tree.insertAt(indexOfCurrentNode, rightNode, nodesInPostOrder)
				nodesToVisit = append(nodesToVisit, rightNode)
			}
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

func NewTree() *Tree {
	return &Tree{
		Root: nil,
	}
}
