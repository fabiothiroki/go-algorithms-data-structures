package main

import "testing"

func TestHello(t *testing.T) {

	r := ReverseString("hello")

	if r != "olleh" {
		t.Errorf("Got: %s, want: %s.", r, "olleh")
	}
}
