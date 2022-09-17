package linkedlists_test

import (
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testHasCycle(testCase *testing.T) {
	var firstNode *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  0,
		Next: nil,
	}

	var secondNode *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  1,
		Next: nil,
	}
	firstNode.Next = secondNode

	var thirdNode *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  2,
		Next: nil,
	}
	secondNode.Next = thirdNode

	var fourthNode *linkedlists.ListNode = &linkedlists.ListNode{
		Val:  3,
		Next: nil,
	}
	thirdNode.Next = fourthNode
	fourthNode.Next = secondNode

	var result bool = linkedlists.HasCycle(firstNode)
	var expectedResult bool = true

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	fourthNode.Next = nil
	result = linkedlists.HasCycle(firstNode)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestHasCycle(testCase *testing.T) {
	testCase.Run("action=test-has-cycle", testHasCycle)
}
