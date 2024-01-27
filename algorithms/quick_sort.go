package main

import "fmt"

/*
Time Complexity:
  Best: Ω(n log(n))
  Average: Θ(n log(n))
  Worst: O(n^2)
Space Complexity:
  O(log(n))
*/
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int

		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func runQuickSort() {
	if ok := testSort(quickSortStart); ok {
		fmt.Println("quick sort success")
	} else {
		fmt.Println("quick sort failure")
	}
}
