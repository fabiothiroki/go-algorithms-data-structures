package strstr

// StrStr implements the respective method of C++
func StrStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	foundIndex := -1
	j := -1

	for i := range haystack {

		if foundIndex == -1 {
			if haystack[i] == needle[0] {
				if len(needle) == 1 {
					return i
				}

				foundIndex = i
				j = 1
				continue
			} else if len(haystack)-i <= len(needle) {
				return -1
			} else {
				continue
			}
		}

		if foundIndex != -1 {
			if haystack[i] == needle[j] {

				if j == len(needle)-1 {
					return foundIndex
				}
				j++

			} else {
				foundIndex = -1
			}
		}

	}

	return foundIndex
}
