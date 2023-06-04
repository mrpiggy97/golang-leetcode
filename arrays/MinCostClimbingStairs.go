package arrays

func GetMinCost(cost []int, i int) int {
	var selectedNode *int = nil
	if i == len(cost)-2 {
		return cost[i]
	}
	var indexChecker map[int]int = make(map[int]int)
	for index := i; index < len(cost); index++ {
		if index == i {
			indexChecker[index] = cost[index]
		}
		if index+1 < len(cost) {
			currentVal, exists := indexChecker[index+1]
			var newVal int = indexChecker[index] + cost[index+1]
			if !exists {
				indexChecker[index+1] = newVal
			}
			if exists && newVal < currentVal {
				indexChecker[index+1] = newVal
			}
			if index+1 == len(cost)-1 || index+1 == len(cost)-2 {
				if selectedNode == nil {
					selectedNode = new(int)
					*selectedNode = indexChecker[index+1]
				}
				if selectedNode != nil && indexChecker[index+1] < *selectedNode {
					*selectedNode = indexChecker[index+1]
				}
			}
		}
		if index+2 < len(cost) {
			currentVal, exists := indexChecker[index+2]
			var newVal int = indexChecker[index] + cost[index+2]
			if !exists {
				indexChecker[index+2] = newVal
			}
			if exists && newVal < currentVal {
				indexChecker[index+2] = newVal
			}
			if index+2 == len(cost)-1 || index+2 == len(cost)-2 {
				if selectedNode == nil {
					selectedNode = new(int)
					*selectedNode = indexChecker[index+2]
				}
				if selectedNode != nil {
					if indexChecker[index+2] < *selectedNode {
						*selectedNode = indexChecker[index+2]
					}
				}
			}
		}
	}
	return *selectedNode
}

func MinCostClimbingStairs(cost []int) int {
	var first int = GetMinCost(cost, 0)
	var second int = GetMinCost(cost, 1)
	if first < second {
		return first
	}
	return second
}
