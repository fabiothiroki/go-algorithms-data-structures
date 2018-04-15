package firstuniqchar

import "testing"

func TestFirstChar(t *testing.T) {

	r := FirstUniqChar("leetcode")

	if r != 0 {
		t.Errorf("Got: %d, want: %d.", r, 0)
	}
}

func TestThirdChar(t *testing.T) {

	r := FirstUniqChar("loveleetcode")

	if r != 2 {
		t.Errorf("Got: %d, want: %d.", r, 2)
	}
}
