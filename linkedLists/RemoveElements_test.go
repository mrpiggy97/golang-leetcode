package linkedlists_test

import (
	"fmt"
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testRemoveItemsWithValues(testCase *testing.T) {
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

	var nodeSlice []*linkedlists.ListNode = linkedlists.GetNodesAsSlice(node1)
	nodeSlice = linkedlists.RemoveItemsWithValue(nodeSlice, 10)
	var expectedLength int = 3
	if len(nodeSlice) != expectedLength {
		testCase.Errorf("expected length of nodeSlice to be %d, got %d", expectedLength, len(nodeSlice))
	}
	for _, node := range nodeSlice {
		fmt.Println(node.Val)
	}
}

func testRemoveElements(testCase *testing.T) {
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

	var result *linkedlists.ListNode = linkedlists.RemoveElements(node1, 10)
	var resultAsSlices []*linkedlists.ListNode = linkedlists.GetNodesAsSlice(result)
	if len(resultAsSlices) != 3 {
		testCase.Errorf("expected length of resultAsSlices to be %d, got %d", 3, len(resultAsSlices))
	}
}

func TestRemoveElements(testCase *testing.T) {
	testCase.Run("action=test-remove-items-with-value", testRemoveItemsWithValues)
	testCase.Run("action=test-remove-elements", testRemoveElements)
}
