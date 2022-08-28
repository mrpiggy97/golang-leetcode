package main

import (
	"fmt"

	"github.com/mrpiggy97/golang-leetcode/trees"
)

func main() {
	var myTree *trees.TreeNode = trees.NewTreeNode()
	myTree.Insert(20)
	myTree.Insert(19)
	myTree.Insert(11)
	myTree.Insert(7)
	myTree.Insert(12)
	myTree.Insert(4)
	myTree.Insert(8)
	myTree.Insert(2)
	myTree.Insert(9)
	var result []int = myTree.InOrderTraverse()
	fmt.Println("result ", result[0])
}
