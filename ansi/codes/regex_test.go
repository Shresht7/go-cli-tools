package codes_test

import (
	"testing"

	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

func TestStrip(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello, \033[31mWorld!\033[0m",
			expected: "Hello, World!",
		},
		{
			input:    "This is a \033[1;4;32mtest\033[0m",
			expected: "This is a test",
		},
	}

	for _, test := range tests {
		result := codes.Strip(test.input)
		if result != test.expected {
			t.Errorf("Strip(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
