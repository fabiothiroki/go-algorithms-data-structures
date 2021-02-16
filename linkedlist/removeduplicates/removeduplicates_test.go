package removeduplicates

import "testing"

func TestNilList(t *testing.T) {
	r := RemoveDuplicates(nil)
	if r != nil {
		t.Errorf("Got: %p, want: nil.", r)
	}
}

func TestDuplicatesList(t *testing.T) {
	n2 := &Node{
		value: 2,
		next:  nil,
	}

	n1 := &Node{
		value: 2,
		next:  n2,
	}

	r := RemoveDuplicates(n1)

	m := CountElements(r)
	if m[2] != 1 {
		t.Errorf("Got: %d, want: 1.", m[2])
	}
}
