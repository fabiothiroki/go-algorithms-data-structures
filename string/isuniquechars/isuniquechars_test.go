package isuniquechars

import "testing"

func TestFalse(t *testing.T) {

	r := isUniqueChar("leetcode")

	if r != false {
		t.Errorf("Got: %t, want: %t.", r, false)
	}
}

func TestTrue(t *testing.T) {

	r := isUniqueChar("letcod")

	if r != true {
		t.Errorf("Got: %t, want: %t.", r, true)
	}
}
