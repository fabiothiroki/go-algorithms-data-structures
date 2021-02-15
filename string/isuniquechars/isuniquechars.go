package isuniquechars

func isUniqueChar(s string) bool {
	m := make(map[string]bool)

	for _, c := range s {
		if m[string(c)] {
			return false
		}

		m[string(c)] = true
	}

	return true
}
