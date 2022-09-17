package linkedlists_test

import (
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testSortLowToHigh(testCase *testing.T) {
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
	var nodes []*linkedlists.ListNode = []*linkedlists.ListNode{node1, node2, node3, node4}
	nodes = linkedlists.SortLowToHigh(nodes)
	var expectedResult []*linkedlists.ListNode = []*linkedlists.ListNode{node2, node4, node3, node1}
	for index, node := range nodes {
		if node.Val != expectedResult[index].Val {
			testCase.Errorf("expected %d to be %d", node.Val, expectedResult[index].Val)
		}
	}
}

func testMergeLists(testCase *testing.T) {
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

	var mainNode *linkedlists.ListNode = linkedlists.MergeLists(node1, node3)
	var expectedResults []int = []int{1, 6, 10, 20}
	var currentNode *linkedlists.ListNode = mainNode

	for _, value := range expectedResults {
		if currentNode.Val != value {
			testCase.Errorf("expected %d to be %d", currentNode.Val, value)
		}
		currentNode = currentNode.Next
	}
}

func TestMergeLists(testCase *testing.T) {
	testCase.Run("action=test-sort-low-to-high", testSortLowToHigh)
	testCase.Run("action=test-merge-lists", testMergeLists)
}
