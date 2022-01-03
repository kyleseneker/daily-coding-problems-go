package main

import "fmt"

func main() {
	var input = []int{3, 4, 9, 6, 1}
	var counter []int

	// populate "counter" slice with zeros
	for i := 0; i < len(input); i++ {
		counter = append(counter, 0)
	}

	// iterate over "input" slice from left to right
	for i := 0; i < len(input); i++ {
		// iterate over all elements on right side of position "i"
		for j := i + 1; j < len(input); j++ {
			// if value at position "j" is less than value at position "i"
			// increment value of counter slice at position "i"
			if input[j] < input[i] {
				counter[i]++
			}
		}
	}

	fmt.Println(counter)
}
