package linkedlists

func HasCycle(node *ListNode) bool {
	var nodeMap map[*ListNode]int = make(map[*ListNode]int)
	var currentNode *ListNode = node
	var currentPos int = 0
	for currentNode != nil {
		_, exists := nodeMap[currentNode]
		if !exists {
			nodeMap[currentNode] = currentPos
		}
		if exists {
			return true
		}
		currentNode = currentNode.Next
		currentPos = currentPos + 1
	}
	return false
}
