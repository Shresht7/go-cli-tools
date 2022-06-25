package helpers

import (
	"testing"
)

func WrapWhitespace(s string) string      { return " " + s + " " }
func WrapDash(s string) string            { return "-" + s + "-" }
func WrapAngularBrackets(s string) string { return "<" + s + ">" }

func TestPipe(t *testing.T) {
	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc:     "should compose the given functions",
			actual:   Compose(WrapAngularBrackets, WrapDash, WrapWhitespace)("Hello Go"),
			expected: "<- Hello Go ->",
		},
		{
			desc:     "should pipe the given functions",
			actual:   Compose(WrapAngularBrackets, WrapDash, WrapWhitespace)("Hello Go"),
			expected: " -<Hello Go>- ",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actual != tC.expected {
				t.Errorf("want:\t%s\ngot:\t%s\n", tC.expected, tC.actual)
			}
		})
	}
}
