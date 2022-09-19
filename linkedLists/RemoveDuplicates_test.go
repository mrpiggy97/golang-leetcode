package linkedlists_test

import (
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testRemoveDuplicates(testCase *testing.T) {
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
	node2.Next = node3

	var node4 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  6,
		Next: nil,
	}
	node3.Next = node4

	var node5 *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  10,
		Next: nil,
	}
	node5.Next = nil
	var result *linkedlists.ListNode = linkedlists.RemoveDuplicates(node1)
	var expectedValues []int = []int{20, 1, 10, 6}
	var currentNode *linkedlists.ListNode = result
	var currentIndex int = 0
	for currentNode != nil {
		if currentNode.Val != expectedValues[currentIndex] {
			testCase.Errorf("expected node value to be %d, got %d", expectedValues[currentIndex], currentNode.Val)
		}
		currentNode = currentNode.Next
		currentIndex = currentIndex + 1
	}
}

func TestRemoveDuplicates(testCase *testing.T) {
	testCase.Run("action=test-remove-duplicates", testRemoveDuplicates)
}
