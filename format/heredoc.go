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

// Checks if the given rune is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// determineMinIndentation determines the minimum indentation level of the given lines.
func determineMinIndentation(lines []string) int {
	// Initialize minimum indentation level
	min := 0

	// Iterate over lines and determine indentation level
	for _, line := range lines {
		indentationLevel := 0 // number of spaces at the beginning of the line
		// Determine indentation level
		for _, r := range line {
			if isSpace(r) {
				indentationLevel++
			} else {
				break
			}
		}

		// Update minimum indentation level
		if indentationLevel < min || min == 0 {
			min = indentationLevel
		}
	}

	// Return minimum indentation level
	return min
}
