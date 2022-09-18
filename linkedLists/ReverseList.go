package linkedlists

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var nodeSlice []*ListNode = GetNodesAsSlice(head)
	var endIndex int = len(nodeSlice) - 1
	var reversedList []*ListNode = []*ListNode{}
	for index := endIndex; index >= 0; index-- {
		var currentNode *ListNode = nodeSlice[index]
		reversedList = append(reversedList, currentNode)
	}
	var currentNode *ListNode = nil
	for index, node := range reversedList {
		if index == 0 {
			currentNode = node
			currentNode.Next = nil
		} else {
			currentNode.Next = node
			currentNode = node
			if index == len(reversedList)-1 {
				currentNode.Next = nil
			}
		}
	}
	return reversedList[0]
}
