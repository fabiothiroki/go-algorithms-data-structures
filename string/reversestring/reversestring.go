package main

import (
	"bytes"
)

// ReverseString prints a string reversed
func ReverseString(s string) string {
	var buffer bytes.Buffer
	i := len(s) - 1

	for i >= 0 {
		buffer.WriteByte(s[i])
		i--
	}

	return buffer.String()
}
