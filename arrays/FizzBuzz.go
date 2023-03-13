package arrays

import "fmt"

func FizzBuzz(n int) []string {
	var result []string = []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		}
		if i%3 == 0 && i%5 != 0 {
			result = append(result, "Fizz")
		}
		if i%3 != 0 && i%5 == 0 {
			result = append(result, "Buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}
