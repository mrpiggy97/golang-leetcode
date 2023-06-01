package arrays

type CostNode struct {
	Left  *CostNode
	Right *CostNode
	Val   int
	Index int
}

func GetMinCost(cost []int, index int) int {
	var rootNode *CostNode = &CostNode{
		Left:  nil,
		Right: nil,
		Val:   cost[index],
		Index: index,
	}
	var nodesToVisit []*CostNode = []*CostNode{rootNode}
	var lowestCost int = 1000
	for len(nodesToVisit) > 0 {
		// form nodes
		var newNodesToVisit []*CostNode = []*CostNode{}
		for _, node := range nodesToVisit {
			if node.Index+1 < len(cost) {
				node.Left = &CostNode{
					Left:  nil,
					Right: nil,
					Val:   node.Val + cost[node.Index+1],
					Index: node.Index + 1,
				}
				newNodesToVisit = append(newNodesToVisit, node.Left)
			}
			if node.Index+2 < len(cost) {
				node.Right = &CostNode{
					Left:  nil,
					Right: nil,
					Val:   node.Val + cost[node.Index+2],
					Index: node.Index + 2,
				}
				newNodesToVisit = append(newNodesToVisit, node.Right)
			}

			if node.Index+1 >= len(cost) || node.Index+2 >= len(cost) {
				if node.Val < lowestCost {
					lowestCost = node.Val
				}
			}
		}
		nodesToVisit = newNodesToVisit
	}
	return lowestCost
}

func MinCostClimbingStairs(cost []int) int {
	var firstCost int = GetMinCost(cost, 0)
	var secondCost int = GetMinCost(cost, 1)
	if firstCost < secondCost {
		return firstCost
	}
	return secondCost
}
