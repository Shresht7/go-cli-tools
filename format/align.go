package format

import (
	"math"
	"strings"

	"github.com/Shresht7/go-cli-tools/helpers"
)

type AlignOptions struct {
	Mode  string
	Split string
	Pad   string
	Width int
}

//	If option values have not been assigned, set them to defaults
func (opts *AlignOptions) AssignValues() {

	//	Split on new lines by default
	if opts.Split == "" {
		opts.Split = "\n"
	}

	//	Use whitespace for padding by default
	if opts.Pad == "" {
		opts.Pad = " "
	}

	switch opts.Mode {
	case "Left", "Center", "Right":
	default:
		opts.Mode = "Center"
	}

}

func Align(s string, opts *AlignOptions) string {

	opts.AssignValues()

	//	Maximum width of the text to display
	maxWidth := opts.Width

	strSlice := strings.Split(s, opts.Split)

	//	Determine maxWidth
	tupleSlice := Map(strSlice, func(split string, _ int) *StringAndWidth {
		width := helpers.StringWidth(split)
		maxWidth = int(math.Max(float64(width), float64(maxWidth)))
		return &StringAndWidth{split: split, width: width}
	})

	assembledSlice := Map(tupleSlice, func(value *StringAndWidth, _ int) string {
		space := maxWidth - value.width
		return applyPadding(value.split, opts, space)
	})

	newStr := strings.Join(assembledSlice, opts.Split)

	return newStr
}

type StringAndWidth struct {
	split string
	width int
}

func Map[T, R any](slice []T, cb func(value T, index int) R) []R {
	newSlice := make([]R, len(slice))
	for idx, value := range slice {
		newSlice[idx] = cb(value, idx)
	}
	return newSlice
}

func applyPadding(s string, opts *AlignOptions, space int) string {
	switch opts.Mode {
	case "Left":
		return PadRight(s, opts.Pad, space)
	case "Center":
		half := int(math.Floor(float64(space) / 2.0))
		if space%2 == 0 {
			return Pad(s, opts.Pad, half)
		} else {
			return strings.Repeat(opts.Pad, half) + s + strings.Repeat(opts.Pad, half+1)
		}
	case "Right":
		return PadLeft(s, opts.Pad, space)
	}
	return s
}
