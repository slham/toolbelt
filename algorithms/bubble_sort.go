package main

import "fmt"

/*
Time Complexity:
  Best: Ω(n)
  Average: Θ(n^2)
  Worst: O(n^2)
Space Complexity:
  O(1)
*/
func bubbleSort(input []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}

	return input
}

func runBubbleSort() {
	if ok := testSort(bubbleSort); ok {
		fmt.Println("bubble sort success")
	} else {
		fmt.Println("bubble sort failure")
	}
}
