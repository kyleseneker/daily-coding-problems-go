package main

import "fmt"

func main() {
	var n = 20
	var outputs []int
	for i := 1; len(outputs) < n; i++ {
		if isRegularNumber(i) {
			outputs = append(outputs, i)
		}
	}
	fmt.Println(outputs)
}

// isRegularNumber determines if a given number x is a "regular" number
// a regular number is a whole number whose only prime divisors are 2, 3, and 5.
func isRegularNumber(x int) bool {
	if x == 1 {
		return true
	}
	if x%2 == 0 {
		return isRegularNumber(x / 2)
	}
	if x%3 == 0 {
		return isRegularNumber(x / 3)
	}
	if x%5 == 0 {
		return isRegularNumber(x / 5)
	}
	return false
}
