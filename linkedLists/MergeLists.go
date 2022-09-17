package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveItemFromSlice(slice []*ListNode, index int) []*ListNode {
	var newSlice []*ListNode = []*ListNode{}
	for currentIndex, value := range slice {
		if currentIndex != index {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func getLowestValue(slice []*ListNode) (lowestValue *ListNode, index int) {
	index = 0
	lowestValue = slice[index]
	for i := index; i < len(slice); i++ {
		var currentValue *ListNode = slice[i]
		if currentValue.Val < lowestValue.Val {
			lowestValue = currentValue
			index = i
		}
	}
	return lowestValue, index
}

func SortLowToHigh(nodeSlice []*ListNode) []*ListNode {
	var sliceCopy []*ListNode = make([]*ListNode, len(nodeSlice))
	copy(sliceCopy, nodeSlice)
	var nodesLowToHigh []*ListNode = []*ListNode{}
	var nodeSliceLength int = len(sliceCopy)

	for nodeSliceLength > 0 {
		currentValue, currentIndex := getLowestValue(sliceCopy)
		nodesLowToHigh = append(nodesLowToHigh, currentValue)
		sliceCopy = RemoveItemFromSlice(sliceCopy, currentIndex)
		nodeSliceLength = len(sliceCopy)
	}
	return nodesLowToHigh
}

func GetNodes(node *ListNode) []*ListNode {
	var slice []*ListNode = []*ListNode{}
	var currentNode *ListNode = node
	for currentNode != nil {
		slice = append(slice, currentNode)
		currentNode = currentNode.Next
	}
	return slice
}

func MergeLists(node1 *ListNode, node2 *ListNode) *ListNode {
	var nodeSlice []*ListNode = GetNodes(node1)
	nodeSlice = append(nodeSlice, GetNodes(node2)...)
	if len(nodeSlice) == 0 {
		return nil
	}
	nodeSlice = SortLowToHigh(nodeSlice)
	var currentNode *ListNode = nil
	for index, node := range nodeSlice {
		if index == 0 {
			currentNode = node
		} else {
			currentNode.Next = node
			currentNode = node
		}
	}
	return nodeSlice[0]
}
