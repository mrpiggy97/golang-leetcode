package trees

func SumOfLeftLeafs(root *TreeNode) int {
	var nodesToVisit []*TreeNode = []*TreeNode{root}
	var sum int = 0
	if root == nil {
		return 0
	}
	for i := 0; i < len(nodesToVisit); i++ {
		var currentNode *TreeNode = nodesToVisit[i]
		if currentNode.Left != nil {
			nodesToVisit = append(nodesToVisit, currentNode.Left)
			if currentNode.Left.Left == nil && currentNode.Left.Right == nil {
				sum += currentNode.Left.Val
			}
		}
		if currentNode.Right != nil {
			nodesToVisit = append(nodesToVisit, currentNode.Right)
		}
	}
	return sum
}
