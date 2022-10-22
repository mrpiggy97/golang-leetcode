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
	var stopSearching bool = false
	for index := range indexes {
		if stopSearching {
			break
		}
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

func SendIndexes(indexChan chan<- int, secondIndexes chan<- int, indexesSorted []int) {
	for index := range indexesSorted {
		if index%2 != 0 {
			indexChan <- index
		} else {
			secondIndexes <- index
		}
	}
	close(indexChan)
	close(secondIndexes)
}

func MaxProfit(prices []int) int {
	var indexesHighToLow []int = GetIndexesHighToLow(prices)
	var profit int = 0
	var indexChan chan int = make(chan int, len(indexesHighToLow))
	var index2Chan chan int = make(chan int, len(indexesHighToLow))
	var results1 chan int = make(chan int, len(indexesHighToLow))
	var results2 chan int = make(chan int, len(indexesHighToLow))
	var results3 chan int = make(chan int, len(indexesHighToLow))
	var results4 chan int = make(chan int, len(indexesHighToLow))
	var results5 chan int = make(chan int, len(indexesHighToLow))
	var results6 chan int = make(chan int, len(indexesHighToLow))
	var results7 chan int = make(chan int, len(indexesHighToLow))
	var results8 chan int = make(chan int, len(indexesHighToLow))
	var results9 chan int = make(chan int, len(indexesHighToLow))
	var results10 chan int = make(chan int, len(indexesHighToLow))

	go SendIndexes(indexChan, index2Chan, indexesHighToLow)
	go ProfitWorker(prices, indexesHighToLow, indexChan, results1)
	go ProfitWorker(prices, indexesHighToLow, indexChan, results2)
	go ProfitWorker(prices, indexesHighToLow, indexChan, results3)
	go ProfitWorker(prices, indexesHighToLow, indexChan, results4)
	go ProfitWorker(prices, indexesHighToLow, indexChan, results5)
	go ProfitWorker(prices, indexesHighToLow, index2Chan, results6)
	go ProfitWorker(prices, indexesHighToLow, index2Chan, results7)
	go ProfitWorker(prices, indexesHighToLow, index2Chan, results8)
	go ProfitWorker(prices, indexesHighToLow, index2Chan, results9)
	go ProfitWorker(prices, indexesHighToLow, index2Chan, results10)

	var results1Open bool = true
	var results2Open bool = true
	var results3Open bool = true
	var results4Open bool = true
	var results5Open bool = true
	var results6Open bool = true
	var results7Open bool = true
	var results8Open bool = true
	var results9Open bool = true
	var results10Open bool = true

	for results1Open || results2Open || results3Open || results4Open || results5Open || results6Open || results7Open || results8Open || results9Open || results10Open {
		select {
		case newProfit, ok := <-results1:
			if !ok {
				results1Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results2:
			if !ok {
				results2Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results3:
			if !ok {
				results3Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results4:
			if !ok {
				results4Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results5:
			if !ok {
				results5Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results6:
			if !ok {
				results6Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results7:
			if !ok {
				results7Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results8:
			if !ok {
				results8Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results9:
			if !ok {
				results9Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		case newProfit, ok := <-results10:
			if !ok {
				results10Open = false
			} else {
				if newProfit > profit {
					profit = newProfit
				}
			}
		}
	}

	return profit
}
