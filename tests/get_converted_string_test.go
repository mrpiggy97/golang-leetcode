package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func TestGetConvertedStringWithTimeout(testCase *testing.T) {
	///this test was made to see how we can skip a test if
	//it is taking too much time
	var testTimer *time.Timer = time.NewTimer(time.Millisecond * 500)
	var channel chan string = make(chan string, 1)
	var stringToConvert string = "faBIAN Is THE Boss"
	var expectedResult string = "Fabian Is The Boss"
	var result string = stringMethods.GetConvertedString(stringToConvert, channel)

	select {
	case <-testTimer.C:
		//timeout has been reached
		testCase.Skip("test timeout has been reached")
	case <-channel:
		if result != expectedResult {
			var errorMessage string = fmt.Sprintf("expected %v to be %v", result, expectedResult)
			testCase.Error(errorMessage)
		}
	}
}
