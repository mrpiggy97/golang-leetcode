package trees

func insert(root *TreeNode, val int) {
	var currentNode *TreeNode = root
	for currentNode != nil {
		if val < currentNode.Val {
			if currentNode.Left != nil {
				currentNode = currentNode.Left
			} else {
				currentNode.Left = &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   val,
				}
				break
			}
		}
		if val > currentNode.Val {
			if currentNode.Right != nil {
				currentNode = currentNode.Right
			} else {
				currentNode.Right = &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   val,
				}
				break
			}
		}
	}
}

func TrimTree(root *TreeNode, low int, high int) *TreeNode {
	var nodesToVisit []*TreeNode = []*TreeNode{root}
	var newHead *TreeNode = nil
	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *TreeNode = nodesToVisit[i]
		if newHead == nil {
			if currentNode.Val >= low && currentNode.Val <= high {
				newHead = &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   currentNode.Val,
				}
			}
		} else {
			if currentNode.Val >= low && currentNode.Val <= high {
				insert(newHead, currentNode.Val)
			}
		}
		if currentNode.Left != nil {
			nodesToVisit = append(nodesToVisit, currentNode.Left)
		}
		if currentNode.Right != nil {
			nodesToVisit = append(nodesToVisit, currentNode.Right)
		}
	}
	return newHead
}
