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

func getLowestValueBeforeEndIndex(endIndex int, priceSlice []int) int {
	if endIndex == 0 {
		return -1
	}
	var lowestValue int = priceSlice[endIndex]
	for index, value := range priceSlice {
		if index == endIndex {
			break
		}
		if value < lowestValue {
			lowestValue = value
		}
	}
	return lowestValue
}

func MaxProfit(prices []int) int {
	var indexes []int = getIndexes(prices)
	var profitsHighToLow []int = make([]int, len(indexes))
	copy(profitsHighToLow, indexes)
	sort.Slice(profitsHighToLow, func(firstIndex, secondIndex int) bool {
		var indexVal1 int = profitsHighToLow[firstIndex]
		var indexVal2 int = profitsHighToLow[secondIndex]
		var val1 int = prices[indexVal1]
		var val2 int = prices[indexVal2]
		return val1 > val2
	})

	var highestProfit int = 0
	var currentProfit int = 0
	for _, highIndex := range profitsHighToLow {
		var highValue int = prices[highIndex]
		var lowestValue int = getLowestValueBeforeEndIndex(highIndex, prices)
		if lowestValue > -1 {
			currentProfit = highValue - lowestValue
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
