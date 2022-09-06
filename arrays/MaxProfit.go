package arrays

import (
	"sort"
)

func getIndexes(prices []int) []int {
	var indexes []int = []int{}
	for index := range prices {
		indexes = append(indexes, index)
	}
	return indexes
}

func MaxProfit(prices []int) int {
	var indexes []int = getIndexes(prices)
	var profitsLowToHigh []int = make([]int, len(indexes))
	copy(profitsLowToHigh, indexes)
	var profitsHighToLow []int = make([]int, len(indexes))
	copy(profitsHighToLow, indexes)
	sort.Slice(profitsLowToHigh, func(firstIndex, secondIndex int) bool {
		var indexVal1 int = profitsLowToHigh[firstIndex]
		var indexVal2 int = profitsLowToHigh[secondIndex]
		var val1 int = prices[indexVal1]
		var val2 int = prices[indexVal2]
		return val1 < val2
	})
	sort.Slice(profitsHighToLow, func(firstIndex, secondIndex int) bool {
		var indexVal1 int = profitsHighToLow[firstIndex]
		var indexVal2 int = profitsHighToLow[secondIndex]
		var val1 int = prices[indexVal1]
		var val2 int = prices[indexVal2]
		return val1 > val2
	})

	var highestProfit int = 0
	var currentProfit int = 0
	for index := range profitsHighToLow {
		var highIndex int = profitsHighToLow[index]
		var highValue int = prices[highIndex]
		for i := range profitsLowToHigh {
			var lowIndex int = profitsLowToHigh[i]
			var lowValue int = prices[lowIndex]
			if lowIndex < highIndex && highValue > lowValue {
				currentProfit = highValue - lowValue
				break
			}
		}
		if currentProfit > highestProfit {
			highestProfit = currentProfit
		}
		if currentProfit < highestProfit {
			break
		}
	}
	return highestProfit
}
