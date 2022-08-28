package trees

import (
	"fmt"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

type Node struct {
	Left      *Node
	Right     *Node
	Value     int
	IsVisited bool
}

func NewNode(value int) *Node {
	return &Node{
		Value:     value,
		Right:     nil,
		Left:      nil,
		IsVisited: false,
	}
}

type TreeNode struct {
	Root         *Node
	nodesToVisit []int
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
	var nodesInOrder *arrays.Array[int] = arrays.NewArray[int]()
	nodesInOrder.Append(&node.Value)

	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *Node = nodesToVisit[i]
		var currentIndex int = *nodesInOrder.GetIndexOfValue(currentNode.Value)
		var leftNode *Node = currentNode.Left
		var rightNode *Node = currentNode.Right

		if leftNode != nil {
			nodesToVisit = append(nodesToVisit, leftNode)
			nodesInOrder.InsertAt(currentIndex, &leftNode.Value)
			currentIndex = *nodesInOrder.GetIndexOfValue(currentNode.Value)
		}
		if rightNode != nil {
			nodesToVisit = append(nodesToVisit, rightNode)
			currentIndex = currentIndex + 1
			if currentIndex == *nodesInOrder.Length {
				nodesInOrder.Append(&rightNode.Value)
			} else {
				nodesInOrder.InsertAt(currentIndex, &rightNode.Value)
			}
		}
	}

	return nodesInOrder.GetSlice()
}

func (treeNode *TreeNode) InOrderTraverse() []int {
	if treeNode.Root == nil {
		fmt.Println("tree is empty")
		return []int{}
	}

	var nodes []int = treeNode.getNodesInOrder(treeNode.Root.Left)
	nodes = append(nodes, treeNode.Root.Value)
	var rightNodes []int = treeNode.getNodesInOrder(treeNode.Root.Right)
	nodes = append(nodes, rightNodes...)
	return nodes
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Root:         nil,
		nodesToVisit: []int{},
	}
}
