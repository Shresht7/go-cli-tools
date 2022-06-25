package format

import "testing"

func TestAlign(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc:     "should align to the center by default",
			actual:   Align("text-align\ncenter", &AlignOptions{}),
			expected: "text-align\n  center  ",
		},
		{
			desc:     "should handle an odd difference",
			actual:   Align("textalign\ncenter", &AlignOptions{}),
			expected: "textalign\n center  ",
		},
		{
			desc:     "should align to the left",
			actual:   Align("textalign\nleft", &AlignOptions{Mode: "Left"}),
			expected: "textalign\nleft     ",
		},
		{
			desc:     "should align to the right",
			actual:   Align("textalign\nright", &AlignOptions{Mode: "Right"}),
			expected: "textalign\n    right",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			if testCase.actual != testCase.expected {
				t.Errorf("want:\t%s\ngot:\t%s\n", testCase.expected, testCase.actual)
			}
		})
	}

}
