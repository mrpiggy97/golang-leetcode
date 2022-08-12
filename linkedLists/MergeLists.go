package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func placeValuesOnSplice(node *ListNode, sendValue chan<- []int) {
	if node != nil {
		var currentNode *ListNode = node
		var nodeValues []int = []int{}
		var currentIndex int = 0
		for currentNode != nil {
			if len(nodeValues) == 0 {
				nodeValues = append(nodeValues, currentNode.Val)
				currentIndex = currentIndex + 1
				currentNode = currentNode.Next
			} else {
				var nextValue int = currentNode.Val
				var previousValue int = nodeValues[currentIndex-1]
				if nextValue < previousValue {
					nodeValues[currentIndex-1] = nextValue
					nodeValues = append(nodeValues, previousValue)
				} else {
					nodeValues = append(nodeValues, nextValue)
				}
				currentIndex = currentIndex + 1
				currentNode = currentNode.Next
			}
		}
		sendValue <- nodeValues
	}

	if node == nil {
		sendValue <- []int{}
	}
}

func MergeLists(firstNode *ListNode, secondNode *ListNode) *ListNode {
	var valuesChannel chan []int = make(chan []int, 1)
	go placeValuesOnSplice(firstNode, valuesChannel)
	go placeValuesOnSplice(secondNode, valuesChannel)
	var finalValues []int = <-valuesChannel
	var valuesOfSecondNode []int = <-valuesChannel
	finalValues = append(finalValues, valuesOfSecondNode...)
	if len(finalValues) == 0 {
		return nil
	}
	// sort final array
	var length int = 1
	for length != len(finalValues) {
		for index, value := range finalValues {
			if index > 0 {
				var previousValue int = finalValues[index-1]
				if value < previousValue {
					finalValues[index-1] = value
					finalValues[index] = previousValue
				}
			}
		}
		length = length + 1
	}
	var mainNode *ListNode = new(ListNode)
	var currentNode *ListNode = mainNode
	for index, value := range finalValues {
		if index == 0 {
			mainNode.Val = value
			mainNode.Next = nil
		}
		if index > 0 {
			var newNode *ListNode = &ListNode{
				Val:  value,
				Next: nil,
			}
			currentNode.Next = newNode
			currentNode = currentNode.Next
		}
	}
	return mainNode
}
