package findintersection

import "testing"

func TestOne(t *testing.T) {
	input := [2]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}

	r := FindIntersection(input[:])

	if r != "1,4,13" {
		t.Errorf("Got: %s, want: %s.", r, "1,4,3")
	}
}

func TestTwo(t *testing.T) {
	input := [2]string{"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"}

	r := FindIntersection(input[:])

	if r != "1,9,10" {
		t.Errorf("Got: %s, want: %s.", r, "1,9,10")
	}
}
