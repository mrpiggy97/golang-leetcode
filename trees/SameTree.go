package trees

func toSlice(node *TreeNode) []*TreeNode {
	var nodesToVisit []*TreeNode = []*TreeNode{node}
	var i int = 0
	for i < len(nodesToVisit) {
		var currentNode *TreeNode = nodesToVisit[i]
		if currentNode != nil {
			nodesToVisit = append(nodesToVisit, currentNode.Left)
			nodesToVisit = append(nodesToVisit, currentNode.Right)
		}
		i += 1
	}
	return nodesToVisit
}

func SameTree(p *TreeNode, q *TreeNode) bool {
	var nodes1 []*TreeNode = toSlice(p)
	var nodes2 []*TreeNode = toSlice(q)
	if len(nodes1) != len(nodes2) {
		return false
	}
	var length int = len(nodes1)
	for i := 0; i < length; i++ {
		var n1 *TreeNode = nodes1[i]
		var n2 *TreeNode = nodes2[i]
		if n1 != nil && n2 != nil {
			if n1.Val != n2.Val {
				return false
			}
		}
		if n1 == nil && n2 != nil {
			return false
		}
		if n1 != nil && n2 == nil {
			return false
		}
	}
	return true
}
