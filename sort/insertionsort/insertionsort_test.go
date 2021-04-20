package insertionsort

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	r := InsertionSort([]int{9, 8, 7, 6})

	if !reflect.DeepEqual(r, []int{6, 7, 8, 9}) {
		t.Errorf("Got: %v, want: %v.", r, []int{6, 7, 8, 9})
	}
}
