package linkedlists

func RemoveItemsWithValue(slice []*ListNode, value int) []*ListNode {
	var newSlice []*ListNode = []*ListNode{}
	for _, node := range slice {
		if node.Val != value {
			newSlice = append(newSlice, node)
		}
	}
	return newSlice
}

func GetNodesAsSlice(node *ListNode) []*ListNode {
	if node == nil {
		return []*ListNode{}
	}
	var nodes []*ListNode = []*ListNode{}
	var currentNode *ListNode = node
	for currentNode != nil {
		nodes = append(nodes, currentNode)
		currentNode = currentNode.Next
	}
	return nodes
}

func RemoveElements(head *ListNode, val int) *ListNode {
	var nodesSlice []*ListNode = GetNodesAsSlice(head)
	if len(nodesSlice) == 0 {
		return nil
	}
	nodesSlice = RemoveItemsWithValue(nodesSlice, val)
	if len(nodesSlice) == 0 {
		return nil
	}
	var currentNode *ListNode = nil
	for index, node := range nodesSlice {
		if index == 0 {
			currentNode = node
			currentNode.Next = nil
		} else {
			currentNode.Next = node
			currentNode = node
			if index == len(nodesSlice)-1 {
				currentNode.Next = nil
			}
		}
	}
	return nodesSlice[0]
}
