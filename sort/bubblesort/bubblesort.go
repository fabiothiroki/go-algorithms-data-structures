package bubblesort

func swap(array[]int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// BubbleSort:
// Best: Time: O(N) Space: O(1)
// Worst and average: Time: O(N^2) Space: O(1)
func BubbleSort(array []int) []int {
	isSorted := false

	for !isSorted {
		isSorted = true

		for i := 0; i < len(array) - 1; i++ {
			if array[i] > array[i+1] {
				swap(array, i, i+1)
				isSorted = false
			}
		}
	}

	return array
}
