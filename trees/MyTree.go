package trees

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Right: nil,
		Left:  nil,
	}
}

type TreeNode struct {
	Root *Node
}

func (treeNode *TreeNode) getIndexOfValue(value int, slice []int) int {
	var correctIndex int = -1
	for index, val := range slice {
		if val == value {
			correctIndex = index
		}
	}
	return correctIndex
}

func (treeNode *TreeNode) insertAt(index int, newValue int, slice []int) []int {
	var values []int = []int{}
	if index == 0 {
		values = append(values, newValue)
		values = append(values, slice...)
	} else {
		var valuesUpToIndex []int = []int{}
		var valuesFromIndexOnward []int = []int{}
		for i := 0; i < index; i++ {
			valuesUpToIndex = append(valuesUpToIndex, slice[i])
		}
		for a := index; a < len(slice); a++ {
			valuesFromIndexOnward = append(valuesFromIndexOnward, slice[a])
		}
		fmt.Println(slice)
		valuesUpToIndex = append(valuesUpToIndex, newValue)
		values = append(values, valuesUpToIndex...)
		var message string = fmt.Sprintf(" new value %d", newValue)
		fmt.Println(slice, message)
		values = append(values, valuesFromIndexOnward...)
	}
	return values
}

func (treeNode *TreeNode) Insert(value int) {

	if treeNode.Root != nil {
		var currentNode *Node = treeNode.Root
		var stop bool = false
		for !stop {
			if value == currentNode.Value {
				fmt.Println("a node with that value already exists")
			}
			if value > currentNode.Value {
				if currentNode.Right == nil {
					currentNode.Right = NewNode(value)
					stop = true
				}
				if currentNode.Right != nil {
					currentNode = currentNode.Right
				}
			}

			if value < currentNode.Value {
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

	if treeNode.Root == nil {
		treeNode.Root = NewNode(value)
	}
}

func (treeNode *TreeNode) Search(value int) *Node {
	var currentNode *Node = treeNode.Root
	var stop bool = false
	for !stop {
		if currentNode.Value == value {
			stop = true
		}

		if value > currentNode.Value {
			if currentNode.Right == nil {
				fmt.Println("no node with that value exists")
				return nil
			}
			if currentNode.Right != nil {
				currentNode = currentNode.Right
			}
		}

		if value < currentNode.Value {
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

func (treeNode *TreeNode) getNodesInOrder(node *Node) []int {
	var nodesToVisit []*Node = []*Node{node}
	var nodesInOrder []int = []int{node.Value}

	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *Node = nodesToVisit[i]
		var currentIndex int = treeNode.getIndexOfValue(currentNode.Value, nodesInOrder)
		var leftNode *Node = currentNode.Left
		var rightNode *Node = currentNode.Right

		if leftNode != nil {
			nodesToVisit = append(nodesToVisit, leftNode)
			nodesInOrder = treeNode.insertAt(currentIndex, leftNode.Value, nodesInOrder)
			currentIndex = treeNode.getIndexOfValue(currentNode.Value, nodesInOrder)
		}
		if rightNode != nil {
			nodesToVisit = append(nodesToVisit, rightNode)
			currentIndex = currentIndex + 1
			if currentIndex == len(nodesInOrder) {
				nodesInOrder = append(nodesInOrder, rightNode.Value)
			} else {
				nodesInOrder = treeNode.insertAt(currentIndex, rightNode.Value, nodesInOrder)
			}
		}
	}

	return nodesInOrder
}

func (treeNode *TreeNode) InOrderTraverse() []int {
	if treeNode.Root == nil {
		fmt.Println("tree is empty")
		return []int{}
	}

	var nodes []int = []int{}

	if treeNode.Root.Left != nil {
		nodes = treeNode.getNodesInOrder(treeNode.Root.Left)
		nodes = append(nodes, treeNode.Root.Value)
	}

	if treeNode.Root.Right != nil {
		var rightNodes []int = treeNode.getNodesInOrder(treeNode.Root.Right)
		nodes = append(nodes, rightNodes...)
	}
	return nodes
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Root: nil,
	}
}
