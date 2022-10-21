package arrays

import (
	"sort"
)

func GetIndexesHighToLow(prices []int) []int {
	// we don't want to alter prices so we use a copy instead
	var pricesCopy []int = make([]int, len(prices))
	copy(pricesCopy, prices)
	// values will serve to retain the current indexes of prices
	// in pricesCopy
	var values map[int][]int = make(map[int][]int)
	var indexesHighToLow []int = []int{}
	for index, val := range pricesCopy {
		slice, exists := values[val]
		if !exists {
			values[val] = []int{index}
		} else {
			slice = append(slice, index)
			values[val] = slice
		}
	}

	// sort pricesCopy high to low
	sort.SliceStable(pricesCopy, func(first, second int) bool {
		return pricesCopy[first] > pricesCopy[second]
	})
	// loop through the sorted slice, get its original index from values map
	// and append that index to indexesHighToLow
	for _, value := range pricesCopy {
		var indexes []int = values[value]
		indexesHighToLow = append(indexesHighToLow, indexes...)
		delete(values, value)
	}
	return indexesHighToLow
}

func ProfitWorker(stockPrices []int, indexesSorted []int, indexes <-chan int, results chan<- int) {
	for index := range indexes {
		var sellingIndex int = indexesSorted[index]
		var sellingPrice int = stockPrices[sellingIndex]
		for tailIndex := len(indexesSorted) - 1; tailIndex > index; tailIndex-- {
			var buyingIndex int = indexesSorted[tailIndex]
			if sellingIndex > buyingIndex {
				var buyingPrice int = stockPrices[buyingIndex]
				var profit int = sellingPrice - buyingPrice
				results <- profit
				break
			}
		}
	}
	close(results)
}

func MaxProfit(prices []int) int {
	var indexesHighToLow []int = GetIndexesHighToLow(prices)
	var profit int = 0
	var indexChan chan int = make(chan int, 1)
	var firstWorkerChan chan int = make(chan int, 1)
	var secondWorkerChan chan int = make(chan int, 1)
	var thirdWorkerChannel chan int = make(chan int, 1)
	go ProfitWorker(prices, indexesHighToLow, indexChan, firstWorkerChan)
	go ProfitWorker(prices, indexesHighToLow, indexChan, secondWorkerChan)
	go ProfitWorker(prices, indexesHighToLow, indexChan, thirdWorkerChannel)

	for index := range indexesHighToLow {
		indexChan <- index
	}
	close(indexChan)

	var firstChannelOpen bool = true
	var secondChannelOpen bool = true
	var thirdChannelOpen bool = true

	for firstChannelOpen || secondChannelOpen || thirdChannelOpen {
		select {
		case newProfit, channelOpen := <-firstWorkerChan:
			if !channelOpen {
				firstChannelOpen = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, channelOpen := <-secondWorkerChan:
			if !channelOpen {
				secondChannelOpen = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, channelOpen := <-thirdWorkerChannel:
			if !channelOpen {
				thirdChannelOpen = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		}
	}

	return profit
}
