package tests

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testCorrectInstances(testCase *testing.T) {
	//this test should not throw error
	var alphanumericInstances []string = make([]string, 6)
	alphanumericInstances[0] = "Mazinkaiser"
	alphanumericInstances[1] = "PassW0rd"
	alphanumericInstances[2] = "43534h56jmTHHF3k"
	alphanumericInstances[3] = "ASDSD234324DDD"
	alphanumericInstances[4] = "123123213"
	alphanumericInstances[5] = "adasdasdasd"

	for _, stringInstance := range alphanumericInstances {
		var isAlphanumeric bool = stringMethods.Parser(stringInstance)
		if !isAlphanumeric {
			fmt.Printf("value %v triggered test to fail\n", stringInstance)
			testCase.Error("expected isAlphanumeric to be true got ", isAlphanumeric)
		}
	}
}

func testErrorInstances(testCase *testing.T) {
	var errorInstances []string = make([]string, 8)
	errorInstances[0] = "&)))((("
	errorInstances[1] = "__ * __"
	errorInstances[2] = "    "
	errorInstances[3] = ""
	errorInstances[4] = "ciao\n$$_"
	errorInstances[5] = "\n\t\n"
	errorInstances[6] = "hello world_"
	errorInstances[7] = ".*?"

	for index, errorInstance := range errorInstances {
		var isAlphanumeric bool = stringMethods.Parser(errorInstance)
		if isAlphanumeric {
			fmt.Printf("%v at index %v in errorInstances triggered test fail\n", errorInstance, index)
			testCase.Error("expected isAlphanumeric to be false got ", isAlphanumeric)
		}
	}
}

func TestParser(testCase *testing.T) {
	testCase.Run("Instances=true", testCorrectInstances)
	testCase.Run("Instances=false", testErrorInstances)
}
