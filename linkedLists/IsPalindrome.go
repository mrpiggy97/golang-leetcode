package linkedlists

type Node struct {
	Val  int
	Next *Node
}

type BNode struct {
	Val      int
	Next     *BNode
	Previous *BNode
}

// Return a doubly linked list's head
func MakeDoublyLinkedLIst(head *Node) (newHead, newTail *BNode) {
	var bhead *BNode = nil
	var current *Node = head
	var currentBNode *BNode = nil
	var counter int = 0
	for current != nil {
		if counter == 0 {
			bhead = &BNode{
				Previous: nil,
				Next:     nil,
				Val:      current.Val,
			}
			newHead = bhead
			counter += 1
			currentBNode = bhead
			current = current.Next
		} else {
			var newBNode *BNode = &BNode{
				Previous: currentBNode,
				Next:     nil,
				Val:      current.Val,
			}
			currentBNode.Next = newBNode
			currentBNode = currentBNode.Next
			current = current.Next
		}
	}
	newTail = currentBNode
	return newHead, newTail
}

func IsPalindrome(head *Node) bool {
	start, end := MakeDoublyLinkedLIst(head)
	for start != end {
		if start.Val != end.Val {
			return false
		}
		start = start.Next
		end = end.Previous
	}
	return true
}
