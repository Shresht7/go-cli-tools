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
