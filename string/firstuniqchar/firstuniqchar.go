package main

// FirstUniqChar returns the index of the first char that appears one time
func FirstUniqChar(s string) int {
	m := make(map[string]int)

	for _, c := range s {
		m[string(c)]++

	}

	for i, c := range s {
		if m[string(c)] == 1 {
			return i
		}
	}

	return -1
}
