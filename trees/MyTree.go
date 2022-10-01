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

func (treeNode *TreeNode) getIndexOfValue(value *Node, slice []*Node) int {
	var correctIndex int = -1
	for index, val := range slice {
		if val == value {
			correctIndex = index
			break
		}
	}
	return correctIndex
}

func (treeNode *TreeNode) insertAt(index int, newValue *Node, slice []*Node) []*Node {
	var values []*Node = []*Node{}
	if index == 0 {
		values = append(values, newValue)
		values = append(values, slice...)
	} else {
		var valuesUpToIndex []*Node = []*Node{}
		var valuesFromIndexOnward []*Node = []*Node{}
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
	var nodesInOrder []*Node = []*Node{node}

	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *Node = nodesToVisit[i]
		var currentIndex int = treeNode.getIndexOfValue(currentNode, nodesInOrder)
		var leftNode *Node = currentNode.Left
		var rightNode *Node = currentNode.Right

		if leftNode != nil {
			nodesToVisit = append(nodesToVisit, leftNode)
			nodesInOrder = treeNode.insertAt(currentIndex, leftNode, nodesInOrder)
			currentIndex = currentIndex + 1
		}
		if rightNode != nil {
			nodesToVisit = append(nodesToVisit, rightNode)
			if currentIndex == len(nodesInOrder)-1 {
				nodesInOrder = append(nodesInOrder, rightNode)
			} else {
				currentIndex = currentIndex + 1
				nodesInOrder = treeNode.insertAt(currentIndex, rightNode, nodesInOrder)
			}
		}
	}

	var valuesOfNodesInOrder []int = []int{}
	for _, pointer := range nodesInOrder {
		valuesOfNodesInOrder = append(valuesOfNodesInOrder, pointer.Value)
	}
	return valuesOfNodesInOrder
}

func (treeNode *TreeNode) getNodesInPreOrder(node *Node) []int {
	var nodesInPostOrder []*Node = []*Node{node}
	var nodesToVisit []*Node = []*Node{node}
	for index := 0; index < len(nodesToVisit); index++ {
		var currentNode *Node = nodesToVisit[index]
		var indexOfCurrentNode int = treeNode.getIndexOfValue(currentNode, nodesInPostOrder)
		var leftNode *Node = currentNode.Left
		var rightNode *Node = currentNode.Right
		var isEndIndex bool = indexOfCurrentNode == len(nodesInPostOrder)-1
		if leftNode != nil {
			if isEndIndex {
				nodesInPostOrder = append(nodesInPostOrder, leftNode)
				nodesToVisit = append(nodesToVisit, leftNode)
			} else {
				indexOfCurrentNode = indexOfCurrentNode + 1
				nodesInPostOrder = treeNode.insertAt(indexOfCurrentNode, leftNode, nodesInPostOrder)
				nodesToVisit = append(nodesToVisit, leftNode)
			}
		}
		indexOfCurrentNode = indexOfCurrentNode + 1
		if rightNode != nil {
			if isEndIndex {
				nodesInPostOrder = append(nodesInPostOrder, rightNode)
				nodesToVisit = append(nodesToVisit, rightNode)
			} else {
				nodesInPostOrder = treeNode.insertAt(indexOfCurrentNode, rightNode, nodesInPostOrder)
				nodesToVisit = append(nodesToVisit, rightNode)
			}
		}
	}
	var nodes []int = make([]int, len(nodesInPostOrder))
	for index, node := range nodesInPostOrder {
		nodes[index] = node.Value
	}
	return nodes
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
	} else {
		nodes = append(nodes, treeNode.Root.Value)
	}

	if treeNode.Root.Right != nil {
		var rightNodes []int = treeNode.getNodesInOrder(treeNode.Root.Right)
		nodes = append(nodes, rightNodes...)
	}
	return nodes
}

func (treeNode *TreeNode) PreOrderTraverse() []int {
	if treeNode.Root == nil {
		return []int{}
	}
	var nodes []int = []int{treeNode.Root.Value}
	var leftNode *Node = treeNode.Root.Left
	var rightNode *Node = treeNode.Root.Right

	if leftNode != nil {
		nodes = append(nodes, treeNode.getNodesInPreOrder(leftNode)...)
	}
	if rightNode != nil {
		nodes = append(nodes, treeNode.getNodesInPreOrder(rightNode)...)
	}
	return nodes
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Root: nil,
	}
}
