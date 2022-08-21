package trees

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

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Root: nil,
	}
}
