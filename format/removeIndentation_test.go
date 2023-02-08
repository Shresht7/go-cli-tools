package format

import (
	"testing"
)

func TestRemoveIndentation(t *testing.T) {

	// Setup test cases
	cases := []struct {
		in, want []string
		n        int
	}{
		{[]string{}, []string{}, 0},
		{[]string{"Hello, World!"}, []string{"Hello, World!"}, 0},
		{[]string{" Hello, World!"}, []string{"Hello, World!"}, 1},
		{[]string{"  Hello, World!"}, []string{"Hello, World!"}, 2},
		{[]string{"\tHello, World!", "\t\tGO!"}, []string{"Hello, World!", "\tGO!"}, 1},
		{[]string{"\t\tHello, World!", "\tGO!"}, []string{"\tHello, World!", "GO!"}, 1},
	}

	// Run test cases
	for _, c := range cases {
		got := RemoveIndentation(c.n, c.in...)
		if !equal(got, c.want) {
			t.Errorf("removeIndentation(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

// equal returns true if the given slices are equal.
func equal(a, b []string) bool {
	// Return false if slices have different length
	if len(a) != len(b) {
		return false
	}
	// Return false if slices have different elements
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	// Return true otherwise
	return true
}
