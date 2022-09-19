package linkedlists

func RemoveDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var keys []int = []int{}
	var nodesMap map[int]*ListNode = make(map[int]*ListNode)
	var currentNode *ListNode = head
	var justBegun bool = true
	for currentNode != nil {
		_, exists := nodesMap[currentNode.Val]
		if !exists {
			nodesMap[currentNode.Val] = currentNode
			keys = append(keys, currentNode.Val)
			currentNode = currentNode.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	currentNode = nil
	for _, key := range keys {
		if justBegun {
			currentNode = nodesMap[key]
			currentNode.Next = nil
			justBegun = false
		} else {
			currentNode.Next = nodesMap[key]
			currentNode = currentNode.Next
		}
	}
	currentNode.Next = nil
	return nodesMap[keys[0]]
}
