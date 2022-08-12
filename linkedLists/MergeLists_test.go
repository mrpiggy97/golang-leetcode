package linkedlists_test

import (
	"fmt"
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func TestMergeLists(testCase *testing.T) {
	var node1 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  20,
		Next: nil,
	}

	var node2 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  1,
		Next: nil,
	}
	node1.Next = node2

	var node3 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  10,
		Next: nil,
	}

	var node4 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  6,
		Next: nil,
	}

	node3.Next = node4

	var finalNode *linkedlists.ListNode = linkedlists.MergeLists(node1, node3)
	if finalNode.Next.Val != 6 {
		var actualResult string = fmt.Sprintf("%d", finalNode.Next.Val)
		testCase.Errorf("finalNode.next.Val is supposed to be 6, got %s instead", actualResult)
	}
}
