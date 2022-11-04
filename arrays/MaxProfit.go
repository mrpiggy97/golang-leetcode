package arrays

func GetProfitForDay(prices []int, indexesChan <-chan int, results chan<- int) {
	for index := range indexesChan {
		var buyingPrice int = prices[index]
		var biggestVal int = prices[index]
		for _, val := range prices[index:] {
			if val > biggestVal {
				biggestVal = val
			}
		}
		if biggestVal != buyingPrice {
			results <- biggestVal - buyingPrice
		}
	}
	close(results)
}

func SendIndexes(prices []int, first, second chan<- int) {
	for index := range prices {
		if index%2 == 0 {
			first <- index
		} else {
			second <- index
		}
	}
	close(first)
	close(second)
}

func MaxProfit(prices []int) int {
	var profit int = 0
	var indexes1 chan int = make(chan int, len(prices))
	var indexes2 chan int = make(chan int, len(prices))

	var results1 chan int = make(chan int, len(prices))
	var results2 chan int = make(chan int, len(prices))
	var results3 chan int = make(chan int, len(prices))
	var results4 chan int = make(chan int, len(prices))
	go SendIndexes(prices, indexes1, indexes2)
	go GetProfitForDay(prices, indexes1, results1)
	go GetProfitForDay(prices, indexes1, results2)
	go GetProfitForDay(prices, indexes2, results3)
	go GetProfitForDay(prices, indexes2, results4)

	var results1Open bool = true
	var results2Open bool = true
	var results3Open bool = true
	var results4Open bool = true

	for results1Open || results2Open || results3Open || results4Open {
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
		}
	}

	return profit
}
