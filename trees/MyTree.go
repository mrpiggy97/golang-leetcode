package trees

import "fmt"

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

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Root: nil,
	}
}
