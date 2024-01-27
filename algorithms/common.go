package main

func testSort(f func([]int) []int) bool {
	arr := []int{8, 3, 4523, 2, 1, 1, 34, 5, 6}
	expected := []int{1, 1, 2, 3, 5, 6, 8, 34, 4523}
	out := f(arr)
	correct := true
	for i := range out {
		if out[i] != expected[i] {
			correct = false
			break
		}
	}
	return correct
}

func main() {
	runBubbleSort()
	runInsertionSort()
	runMergeSort()
	runQuickSort()
}
