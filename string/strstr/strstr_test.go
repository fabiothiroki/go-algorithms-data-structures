package strstr

import "testing"

func TestHello(t *testing.T) {

	r := StrStr("hello", "ll")

	if r != 2 {
		t.Errorf("Got: %d, want: %d.", r, 2)
	}
}

func TestNotFound(t *testing.T) {

	r := StrStr("aaaaa", "bba")

	if r != -1 {
		t.Errorf("Got: %d, want: %d.", r, -1)
	}
}

func TestReset(t *testing.T) {

	r := StrStr("aaabaaaabb", "abb")

	if r != 7 {
		t.Errorf("Got: %d, want: %d.", r, 7)
	}
}

func TestShortNeedle(t *testing.T) {

	r := StrStr("aaa", "a")

	if r != 0 {
		t.Errorf("Got: %d, want: %d.", r, 0)
	}
}

func TestBigNeedle(t *testing.T) {

	r := StrStr("aaa", "aaaa")

	if r != -1 {
		t.Errorf("Got: %d, want: %d.", r, -1)
	}
}
