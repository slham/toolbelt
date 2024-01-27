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
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

func runInsertionSort() {
	if ok := testSort(insertionSort); ok {
		fmt.Println("insertion sort success")
	} else {
		fmt.Println("insertion sort failure")
	}
}
