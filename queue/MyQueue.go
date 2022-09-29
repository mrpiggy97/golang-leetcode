package queue

type MyNode struct {
	Val  int
	Next *MyNode
}

type MyQueue struct {
	Head   *MyNode
	Bottom *MyNode
	Length int
}

func Constructor() *MyQueue {
	return &MyQueue{
		Head:   nil,
		Bottom: nil,
		Length: 0,
	}
}

func (this *MyQueue) Push(x int) {
	if this.Head == nil {
		var newNode *MyNode = &MyNode{
			Val:  x,
			Next: nil,
		}
		this.Head = newNode
		this.Bottom = newNode
		this.Length = this.Length + 1
	} else {
		var newNode *MyNode = &MyNode{
			Val:  x,
			Next: nil,
		}
		this.Bottom.Next = newNode
		this.Bottom = this.Bottom.Next
		this.Length = this.Length + 1
	}
}

func (this *MyQueue) Pop() int {
	var oldHead *MyNode = this.Head
	if this.Head.Next == nil {
		this.Head = nil
		this.Bottom = nil
		this.Length = 0
		return oldHead.Val
	}
	this.Head = this.Head.Next
	this.Length = this.Length - 1
	return oldHead.Val
}

func (this *MyQueue) Peek() int {
	return this.Head.Val
}

func (this *MyQueue) Empty() bool {
	return this.Length == 0
}
