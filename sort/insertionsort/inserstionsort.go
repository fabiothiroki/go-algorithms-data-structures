package insertionsort

// InsertionSort:
// Average/Worst Time: O(n^2) / Space: O(1)
// Best: time O(n) space O(1)
func InsertionSort(array []int) []int {

	for i := 1; i < len(array)-1; i++ {

		for j := i; j >= 0 && array[j-1] < array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}

	return array
}
