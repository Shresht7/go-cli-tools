package helpers

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

func StringWidth(s string) int {

	s = codes.Strip(s)

	//	TODO: Implementation Pending

	width := 0
	for range s {
		width++
	}

	return width
}
