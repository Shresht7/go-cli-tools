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
	return codes.New(styles["bold"]).Wrap(str)
}

// Makes the string faint
func Faint(str string) string {
	return codes.New(styles["faint"]).Wrap(str)
}

// Makes the string italic
func Italic(str string) string {
	return codes.New(styles["italic"]).Wrap(str)
}

// Makes the string underlined
func Underline(str string) string {
	return codes.New(styles["underline"]).Wrap(str)
}

// Makes the string blink
func Blinking(str string) string {
	return codes.New(styles["blinking"]).Wrap(str)
}

// Inverts the string's colors
func Inverse(str string) string {
	return codes.New(styles["inverse"]).Wrap(str)
}

// Hides the string
func Hidden(str string) string {
	return codes.New(styles["hidden"]).Wrap(str)
}

// Strikethrough a string
func Strikethrough(str string) string {
	return codes.New(styles["strikethrough"]).Wrap(str)
}

// Resets the string
func Reset(str string) string {
	return codes.New(styles["reset"]).Wrap(str)
}
