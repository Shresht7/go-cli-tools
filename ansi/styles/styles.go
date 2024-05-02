package styles

import (
	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

//	======
//	STYLES
//	======

var styles = map[string][]string{
	"bold":          {"1", "22"}, //	21 doesn't work for some reason, 22 does the trick though
	"faint":         {"2", "22"},
	"italic":        {"3", "23"},
	"underline":     {"4", "24"},
	"blinking":      {"5", "25"},
	"inverse":       {"7", "27"},
	"hidden":        {"8", "28"},
	"strikethrough": {"9", "29"},
	"reset":         {"0", "0"},
}

// Makes the string bold
func Bold(str string) string {
	return codes.Wrap(str, styles["bold"])
}

// Makes the string faint
func Faint(str string) string {
	return codes.Wrap(str, styles["faint"])
}

// Makes the string italic
func Italic(str string) string {
	return codes.Wrap(str, styles["italic"])
}

// Makes the string underlined
func Underline(str string) string {
	return codes.Wrap(str, styles["underline"])
}

// Makes the string blink
func Blinking(str string) string {
	return codes.Wrap(str, styles["blinking"])
}

// Inverts the string's colors
func Inverse(str string) string {
	return codes.Wrap(str, styles["inverse"])
}

// Hides the string
func Hidden(str string) string {
	return codes.Wrap(str, styles["hidden"])
}

// Strikethrough a string
func Strikethrough(str string) string {
	return codes.Wrap(str, styles["strikethrough"])
}

// Resets the string
func Reset(str string) string {
	return codes.Wrap(str, styles["reset"])
}
