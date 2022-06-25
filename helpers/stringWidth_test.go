package helpers

import "testing"

func TestStringWidth(t *testing.T) {

	testCases := []struct {
		desc  string
		str   string
		width int
	}{
		{
			desc:  "should return 1 for a single character string",
			str:   "a",
			width: 1,
		},
		{
			desc:  "should return 2 for a string with two characters",
			str:   "ab",
			width: 2,
		},
		{
			desc:  "should return 2 for a string with two characters",
			str:   "🍕",
			width: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			actual := StringWidth(testCase.str)
			if actual != testCase.width {
				t.Errorf("want:\t%d\ngot:\t%d\n", testCase.width, actual)
			}
		})
	}

}
