package helpers

// determineMinIndentation determines the minimum indentation level of the given lines.
func DetermineMinIndentation(lines []string) int {
	// Initialize minimum indentation level
	min := 0

	// Iterate over lines and determine indentation level
	for _, line := range lines {
		indentationLevel := 0 // number of spaces at the beginning of the line
		// Determine indentation level
		for _, r := range line {
			if IsSpace(r) {
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
