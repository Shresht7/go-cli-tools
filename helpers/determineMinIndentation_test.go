package helpers

import (
	"testing"
)

func TestDetermineIndentationLevel(t *testing.T) {

	// Setup test cases
	cases := []struct {
		strLines       []string
		minIndentLevel int
	}{
		{[]string{}, 0},
		{[]string{"Test String..."}, 0},
		{[]string{" Test String..."}, 1},
		{[]string{"  Test String..."}, 2},
		{[]string{"\tTest String...", "\t\tGO!"}, 1},
		{[]string{"\t\tTest String...", "\tGO!"}, 1},
		{[]string{"\t\tTest String...", "GO!"}, 0},
		{[]string{"\t\t\tTest String...", "   GO!"}, 3},
	}

	// Run test cases
	for _, c := range cases {
		got := DetermineMinIndentation(c.strLines)
		if got != c.minIndentLevel {
			t.Errorf("determineMinIndentation(%q) == %d, want %d", c.strLines, got, c.minIndentLevel)
		}
	}

}
