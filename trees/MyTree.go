package trees

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type NodeSum struct {
	Node *TreeNode
	Sum  int
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

func (tree *Tree) insertAtStart(slice []int, newValue int) []int {
	var newValues []int = []int{newValue}
	newValues = append(newValues, slice...)
	return newValues
}

func (tree *Tree) ZigZagLevelTraverse() [][]int {
	if tree.Root == nil {
		return [][]int{}
	}
	var currentLevel []*TreeNode = []*TreeNode{tree.Root}
	var values []int = []int{tree.Root.Val}
	var zigZag [][]int = [][]int{values}
	var startFromEnd bool = true
	var length int = len(currentLevel)
	for length > 0 {
		var newLevel []*TreeNode = []*TreeNode{}
		var newValues []int = []int{}
		for _, node := range currentLevel {
			var leftNode *TreeNode = node.Left
			var rightNode *TreeNode = node.Right

			if leftNode != nil {
				newLevel = append(newLevel, leftNode)
				if startFromEnd {
					newValues = tree.insertAtStart(newValues, leftNode.Val)
				} else {
					newValues = append(newValues, leftNode.Val)
				}
			}
			if rightNode != nil {
				newLevel = append(newLevel, rightNode)
				if startFromEnd {
					newValues = tree.insertAtStart(newValues, rightNode.Val)
				} else {
					newValues = append(newValues, rightNode.Val)
				}
			}
		}
		if startFromEnd {
			startFromEnd = false
		} else {
			startFromEnd = true
		}
		if len(newValues) > 0 {
			zigZag = append(zigZag, newValues)
		}
		currentLevel = newLevel
		length = len(currentLevel)
	}
	return zigZag
}

func (tree *Tree) GetDepth() int {
	if tree.Root == nil {
		return 0
	}
	var currentDepth int = 1
	var currentLevel []*TreeNode = []*TreeNode{tree.Root}
	var length int = len(currentLevel)
	for length > 0 {
		var newLevel []*TreeNode = []*TreeNode{}
		for _, node := range currentLevel {
			var leftNode *TreeNode = node.Left
			var rightNode *TreeNode = node.Right

			if leftNode != nil {
				newLevel = append(newLevel, leftNode)
			}
			if rightNode != nil {
				newLevel = append(newLevel, rightNode)
			}
		}
		if len(newLevel) > 0 {
			currentDepth = currentDepth + 1
		}
		currentLevel = newLevel
		length = len(currentLevel)
	}
	return currentDepth
}

func (tree *Tree) removeVal(slice []*TreeNode, index int) []*TreeNode {
	var newValues []*TreeNode = []*TreeNode{}
	for pointerIndex, pointer := range slice {
		if pointerIndex != index {
			newValues = append(newValues, pointer)
		}
	}
	return newValues
}

func (tree *Tree) LevelIsMirror(level []*TreeNode) bool {
	if len(level)%2 != 0 {
		return false
	}
	var lengthOfSlices int = len(level) / 2
	for lengthOfSlices > 0 {
		fmt.Println(level)
		var endIndex int = len(level) - 1
		if level[0] != nil && level[endIndex] != nil {
			if level[0].Val != level[endIndex].Val {
				return false
			}
		}
		if level[0] == nil && level[endIndex] != nil {
			return false
		}
		if level[0] != nil && level[endIndex] == nil {
			return false
		}
		level = tree.removeVal(level, 0)
		level = tree.removeVal(level, len(level)-1)
		lengthOfSlices = len(level)
	}
	return true
}

func (tree *Tree) IsMirror() bool {
	if tree.Root == nil {
		return false
	}
	var currentLevel []*TreeNode = []*TreeNode{tree.Root}
	var nodeNumber int = 1
	//get level
	for nodeNumber > 0 {
		var newLevel []*TreeNode = []*TreeNode{}
		var newNodeNumber int = 0
		for _, node := range currentLevel {
			if node != nil {
				if node.Left != nil {
					newNodeNumber = nodeNumber + 1
				}
				if node.Right != nil {
					newNodeNumber = nodeNumber + 1
				}
				newLevel = append(newLevel, node.Left)
				newLevel = append(newLevel, node.Right)
			}
		}

		nodeNumber = newNodeNumber

		if nodeNumber > 0 {
			//check if level is a mirror
			if len(newLevel)%2 != 0 && len(newLevel) > 0 {
				return false
			}
			if len(newLevel)%2 == 0 {
				var levelIsValid bool = tree.LevelIsMirror(newLevel)
				if !levelIsValid {
					return false
				}
				currentLevel = newLevel
			}
		}
	}
	return true
}

func (tree *Tree) InvertTree() *TreeNode {
	if tree.Root != nil {
		var nodesToVisit []*TreeNode = []*TreeNode{tree.Root}
		for len(nodesToVisit) > 0 {
			var newNodesToVisit []*TreeNode = []*TreeNode{}
			for _, node := range nodesToVisit {
				var leftNode *TreeNode = node.Left
				var rightNode *TreeNode = node.Right
				node.Left = rightNode
				node.Right = leftNode
				if leftNode != nil {
					newNodesToVisit = append(newNodesToVisit, leftNode)
				}
				if rightNode != nil {
					newNodesToVisit = append(newNodesToVisit, rightNode)
				}
			}
			nodesToVisit = newNodesToVisit
		}
	}
	return tree.Root
}

func (tree *Tree) HasPathSum(targetSum int) bool {
	if tree.Root != nil {
		var rootSum *NodeSum = &NodeSum{
			Node: tree.Root,
			Sum:  tree.Root.Val,
		}
		var nodesToVisit []*NodeSum = []*NodeSum{rootSum}
		for len(nodesToVisit) > 0 {
			var newNodesToVisit []*NodeSum = []*NodeSum{}
			for _, node := range nodesToVisit {
				var currentNode *TreeNode = node.Node
				if currentNode != nil {
					var leftNode *TreeNode = currentNode.Left
					var rightNode *TreeNode = currentNode.Right
					if leftNode == nil && rightNode == nil {
						if node.Sum == targetSum {
							return true
						}
					}
					if leftNode != nil {
						var newNode *NodeSum = &NodeSum{
							Node: leftNode,
							Sum:  node.Sum + leftNode.Val,
						}
						newNodesToVisit = append(newNodesToVisit, newNode)
					}
					if rightNode != nil {
						var newNode *NodeSum = &NodeSum{
							Node: rightNode,
							Sum:  node.Sum + rightNode.Val,
						}
						newNodesToVisit = append(newNodesToVisit, newNode)
					}
				}
			}
			nodesToVisit = newNodesToVisit
		}
	}
	return false
}

func NewTree() *Tree {
	return &Tree{
		Root: nil,
	}
}
