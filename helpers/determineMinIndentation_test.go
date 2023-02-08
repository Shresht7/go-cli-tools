package helpers

import (
	"testing"
)

func TestDetermineIndentationLevel(t *testing.T) {

	// Setup test cases
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{}, 0},
		{[]string{"Hello, World!"}, 0},
		{[]string{" Hello, World!"}, 1},
		{[]string{"  Hello, World!"}, 2},
		{[]string{"\tHello, World!", "\t\tGO!"}, 1},
		{[]string{"\t\tHello, World!", "\tGO!"}, 1},
		{[]string{"\t\tHello, World!", "GO!"}, 0},
	}

	// Run test cases
	for _, c := range cases {
		got := DetermineMinIndentation(c.in)
		if got != c.want {
			t.Errorf("determineMinIndentation(%q) == %d, want %d", c.in, got, c.want)
		}
	}

}
