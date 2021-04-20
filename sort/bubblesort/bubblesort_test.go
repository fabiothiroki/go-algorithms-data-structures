package bubblesort

import ("testing" 
"reflect")

func Test(t *testing.T) {
	r := BubbleSort([]int{8, 5, 2, 9, 5, 6, 3})

	if !reflect.DeepEqual(r, []int{ 2, 3, 5, 5, 6, 8, 9 }) {
		t.Errorf("Got: %v, want: %v.", r, []int{ 2, 3, 5, 5, 6, 8, 9 })
	}
}