package main

import "fmt"

/*
Time Complexity:
  Best: Ω(n log(n))
  Average: Θ(n log(n))
  Worst: O(n log(n))
Space Complexity:
  O(n)
*/
func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func runMergeSort() {
	if ok := testSort(mergeSort); ok {
		fmt.Println("merge sort success")
	} else {
		fmt.Println("merge sort failure")
	}
}
