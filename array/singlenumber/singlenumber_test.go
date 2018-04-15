package singlenumber

import "testing"

func TestOneElement(t *testing.T) {
	r := SingleNumber([]int{1})
	if r != 1 {
		t.Errorf("Got: %d, want: %d.", r, 1)
	}
}

func TestThreeElements(t *testing.T) {
	r := SingleNumber([]int{4, 4, 5})
	if r != 5 {
		t.Errorf("Got: %d, want: %d.", r, 1)
	}
}
