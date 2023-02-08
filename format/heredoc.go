package format

import (
	"fmt"
)

// HereDoc returns a here-document representation of the given string.
// See https://en.m.wikipedia.org/wiki/Here_document for more information.
func HereDoc(s string) string {
	return s
}

// HereDocf returns a here-document (https://en.m.wikipedia.org/wiki/Here_document)
// representation of the given string. The can be formatted using
// the given arguments as in fmt.Sprintf.
func HereDocf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

// removeIndentation removes the given number of spaces from the beginning of each line.
func removeIndentation(lines []string, n int) []string {
	for i, line := range lines {
		if len(line) > n {
			lines[i] = line[n:]
		}
	}
	return lines
}
