package linkedlists_test

import (
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testIsPalindrome(testCase *testing.T) {
	var headNode *linkedlists.Node = &linkedlists.Node{
		Next: nil,
		Val:  1,
	}
	var secondNode *linkedlists.Node = &linkedlists.Node{
		Next: nil,
		Val:  2,
	}
	headNode.Next = secondNode

	var thirdNode *linkedlists.Node = &linkedlists.Node{
		Val:  2,
		Next: nil,
	}
	secondNode.Next = thirdNode
	var fourthNode *linkedlists.Node = &linkedlists.Node{
		Val:  1,
		Next: nil,
	}
	thirdNode.Next = fourthNode
	var expectedResult bool = true
	var result bool = linkedlists.IsPalindrome(headNode)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	fourthNode.Val = 100
	expectedResult = false
	result = linkedlists.IsPalindrome(headNode)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestIsPalindrome(testCase *testing.T) {
	testCase.Run("action=test-linked-list-is-palindrome", testIsPalindrome)
}
