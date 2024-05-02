package codes_test

import (
	"testing"

	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

// Wrap a string with ANSI codes
func TestWrap(t *testing.T) {
	code := codes.New([]string{"31", "0"})

	// Test when isEnabled is true
	codes.Enable()
	expected := "\x1b[31mHello\x1b[0m"
	result := code.Wrap("Hello")
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}

	// Test when isEnabled is false
	codes.Disable()
	expected = "Hello"
	result = code.Wrap("Hello")
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
