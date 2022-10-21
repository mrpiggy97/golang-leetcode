package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testGetIndexesHighToLow(testCase *testing.T) {
	var testSlice []int = []int{1, 100, 5, 19, 99}
	var result []int = arrays.GetIndexesHighToLow(testSlice)
	var expectedResult []int = []int{1, 4, 3, 2, 0}
	for index, val := range result {
		if val != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", val, expectedResult[index])
		}
	}
}

func testProfitWorker(testCase *testing.T) {
	var testPrices []int = []int{2, 5, 99, 1, 32, 100}
	var sortedIndexes []int = arrays.GetIndexesHighToLow(testPrices)
	var resultsChan chan int = make(chan int, 1)
	var indexesChan chan int = make(chan int, 1)
	var profitsChannel chan int = make(chan int, 1)
	go func(results <-chan int, finalResult chan<- int) {
		var profit int = 0
		for newProfit := range results {
			if newProfit > profit {
				profit = newProfit
			}
		}
		finalResult <- profit
		close(finalResult)
	}(resultsChan, profitsChannel)
	go arrays.ProfitWorker(testPrices, sortedIndexes, indexesChan, resultsChan)
	for index := range sortedIndexes {
		indexesChan <- index
	}
	close(indexesChan)
	var result int = <-profitsChannel
	var expectedResult int = 99
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func testMaxProfit(testCase *testing.T) {

	testingArray := []int{4, 7, 1, 2}
	bestProfit := arrays.MaxProfit(testingArray)
	expectedResult := 3
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{7, 1, 5, 3, 6, 4}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 5
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{7, 6, 4, 3, 1}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 0
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{4, 11, 1, 2, 7}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 7
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}
	testingArray = []int{1, 4, 1, 4, 3, 1}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 3

	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}
}

func TestMaxProfit(testCase *testing.T) {
	testCase.Run("action=test-max-profit", testMaxProfit)
	testCase.Run("action=test-profit-worker", testProfitWorker)
	testCase.Run("action=test-get-indexes-high-to-low", testGetIndexesHighToLow)
}
