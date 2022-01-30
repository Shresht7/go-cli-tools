package styles

import (
	"strings"

	"github.com/Shresht7/go-cli-tools/helpers"
)

//	=======
//	PADDING
//	=======

//	Apply padding around the string
func Pad(s string, count int) string {
	space := strings.Repeat(" ", count)
	s = space + s + space
	return s
}

//	Apply padding to the left of the string
func PadLeft(s string, count int) string {
	return strings.Repeat(" ", count) + s
}

//	Apply padding to the right of the string
func PadRight(s string, count int) string {
	return s + strings.Repeat(" ", count)
}

//	======
//	STYLES
//	======

var styles = map[string][2]string{
	"bold":          {"1", "22"},
	"faint":         {"2", "22"},
	"italic":        {"3", "23"},
	"underline":     {"4", "24"},
	"blinking":      {"5", "25"},
	"inverse":       {"7", "27"},
	"hidden":        {"8", "28"},
	"strikethrough": {"9", "29"},
}

//	Makes the string bold
func Bold(str string) string {
	return helpers.Wrap(str, styles["bold"])
}

//	Makes the string faint
func Faint(str string) string {
	return helpers.Wrap(str, styles["faint"])
}

//	Makes the string italic
func Italic(str string) string {
	return helpers.Wrap(str, styles["italic"])
}

//	Makes the string underlined
func Underline(str string) string {
	return helpers.Wrap(str, styles["underline"])
}

//	Makes the string blink
func Blinking(str string) string {
	return helpers.Wrap(str, styles["blinking"])
}

//	Inverts the string's colors
func Inverse(str string) string {
	return helpers.Wrap(str, styles["inverse"])
}

//	Hides the string
func Hidden(str string) string {
	return helpers.Wrap(str, styles["hidden"])
}

//	Strikethrough a string
func Strikethrough(str string) string {
	return helpers.Wrap(str, styles["strikethrough"])
}
