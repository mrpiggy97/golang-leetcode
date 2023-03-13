package trees

// separating logic between inserting and getting value
// to insert into tree helps a lot
func insertIntoTree(head *TreeNode, value int) {
	var currentNode *TreeNode = head
	for currentNode != nil {
		if value > currentNode.Val {
			if currentNode.Right == nil {
				currentNode.Right = &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   value,
				}
				break
			} else {
				currentNode = currentNode.Right
			}
		} else {
			if currentNode.Left == nil {
				currentNode.Left = &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   value,
				}
				break
			} else {
				currentNode = currentNode.Left
			}
		}
	}
}

func SortBinaryTree(nums []int) *TreeNode {
	var head *TreeNode = &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   -1,
	}
	if len(nums) == 1 {
		head.Val = nums[0]
		return head
	}
	var headNodeIndex int = int(len(nums) / 2)
	head.Val = nums[headNodeIndex]
	var firstSlice []int = nums[0:headNodeIndex]
	var slicesToVisit [][]int = [][]int{firstSlice}
	if headNodeIndex+1 < len(nums) {
		var secondSlice []int = nums[headNodeIndex+1:]
		slicesToVisit = append(slicesToVisit, secondSlice)
	}
	for currentIndex := 0; currentIndex < len(slicesToVisit); currentIndex++ {
		var currentSlice []int = slicesToVisit[currentIndex]
		if len(currentSlice) == 1 {
			var nodeValue int = currentSlice[0]
			insertIntoTree(head, nodeValue)
		} else {
			var splitIndex int = int(len(currentSlice) / 2)
			var nodeValue int = currentSlice[splitIndex]
			insertIntoTree(head, nodeValue)
			if splitIndex > 0 {
				var leftSlice []int = currentSlice[0:splitIndex]
				slicesToVisit = append(slicesToVisit, leftSlice)
			}
			if splitIndex+1 < len(currentSlice) {
				var rightSlice []int = currentSlice[splitIndex+1:]
				slicesToVisit = append(slicesToVisit, rightSlice)
			}

		}
	}
	return head
}
