package format

import (
	"testing"
)

func TestHereDoc(t *testing.T) {

	// Setup test cases
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"Hello, World!", "Hello, World!"},
		{`Hello, World!`, `Hello, World!`},
		{`
			Hello, World!
			`,
			"Hello, World!",
		},
		{`
			One
			Two
			Three
				GO!
			`,
			"One\nTwo\nThree\n\tGO!",
		},
	}

	// Run test cases
	for _, c := range cases {
		got := HereDoc(c.in)
		if got != c.want {
			t.Errorf("Doc(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

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
		got := determineMinIndentation(c.in)
		if got != c.want {
			t.Errorf("determineMinIndentation(%q) == %d, want %d", c.in, got, c.want)
		}
	}

}
