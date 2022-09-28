package queue

type Node struct {
	Val   rune
	Index int
	Next  *Node
}

type Queue struct {
	Head   *Node
	Bottom *Node
}

func (queue *Queue) FillQueue(str string) {
	for index, value := range str {
		if index == 0 {
			var newNode *Node = &Node{
				Val:   value,
				Index: index,
				Next:  nil,
			}
			queue.Head = newNode
			queue.Bottom = queue.Head
		} else {
			var newNode *Node = &Node{
				Val:   value,
				Index: index,
				Next:  nil,
			}
			queue.Bottom.Next = newNode
			queue.Bottom = queue.Bottom.Next
		}
	}
}

func RemoveNodes(queue *Queue, openerIndex, closureIndex int) *Queue {
	var currentNode *Node = queue.Head
	var index int = 0
	var newQueue *Queue = &Queue{
		Head:   nil,
		Bottom: nil,
	}
	for currentNode != nil {
		if currentNode.Index == openerIndex || currentNode.Index == closureIndex {
			currentNode = currentNode.Next
			if currentNode == nil {
				break
			}
		}
		if currentNode.Index != openerIndex && currentNode.Index != closureIndex {
			if index == 0 {
				var newNode = &Node{
					Val:   currentNode.Val,
					Index: index,
					Next:  nil,
				}
				newQueue.Head = newNode
				newQueue.Bottom = newNode
			}
			if index > 0 {
				var newNode = &Node{
					Val:   currentNode.Val,
					Index: index,
					Next:  nil,
				}
				newQueue.Bottom.Next = newNode
				newQueue.Bottom = newQueue.Bottom.Next
			}
			index = index + 1
			currentNode = currentNode.Next
		}
	}
	return newQueue
}

func ExpectedValue(opener, closure rune) bool {
	var expectedValue rune
	switch opener {
	case 123:
		expectedValue = 125
	case 40:
		expectedValue = 41
	case 91:
		expectedValue = 93
	}
	return expectedValue == closure
}

func IsOpener(str rune) bool {
	if str == 123 || str == 40 || str == 91 {
		return true
	} else {
		return false
	}
}

func IsValid(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	if len(str) == 0 {
		return false
	}
	if rune(str[0]) != 123 && rune(str[0]) != 40 && rune(str[0]) != 91 {
		return false
	}
	var lastIndex int = len(str) - 1
	if rune(str[lastIndex]) != 125 && rune(str[lastIndex]) != 41 && rune(str[lastIndex]) != 93 {
		return false
	}

	var queue *Queue = &Queue{
		Head:   nil,
		Bottom: nil,
	}
	queue.FillQueue(str)
	var stop bool = false
	var currentNode *Node = queue.Head
	for !stop {
		var nodeIsOpener bool = IsOpener(currentNode.Val)
		if !nodeIsOpener {
			return false
		}
		if nodeIsOpener {
			if currentNode.Next == nil {
				return false
			}
			if currentNode.Next != nil {
				var nextNodeIsOpener bool = IsOpener(currentNode.Next.Val)
				if nextNodeIsOpener {
					currentNode = currentNode.Next
				}
				if !nextNodeIsOpener {
					var closureIsValid bool = ExpectedValue(currentNode.Val, currentNode.Next.Val)
					if !closureIsValid {
						return false
					}
					if closureIsValid {
						queue = RemoveNodes(queue, currentNode.Index, currentNode.Next.Index)
						if queue.Head == nil {
							break
						}
						currentNode = queue.Head
					}
				}
			}
		}
	}

	return true
}
