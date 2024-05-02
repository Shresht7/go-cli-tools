package colors

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

// Base Colors
var baseColors = map[string][2]int{
	"black":   {30, 39},
	"red":     {31, 39},
	"green":   {32, 39},
	"yellow":  {33, 39},
	"blue":    {34, 39},
	"magenta": {35, 39},
	"cyan":    {36, 39},
	"white":   {37, 39},
	"default": {39, 39},
}

// Background  color offset
const BG_OFFSET = 10

// Bright color offset
const BRIGHT_OFFSET = 60

// Foreground Colors
var colors = make(map[string][]string)

// Background Colors
var bgColors = make(map[string][]string)

// Bright Foreground Colors
var brightColors = make(map[string][]string)

// Bright Background Colors
var brightBgColors = make(map[string][]string)

// Formats given integers as a string tuple
func itoaSlice(i ...int) []string {
	var s []string
	for _, v := range i {
		s = append(s, strconv.Itoa(v))
	}
	return s
}

// Capitalizes a string
func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func init() {
	for key, tuple := range baseColors {
		colors[key] = itoaSlice(tuple[0], tuple[1])
		bgColors["bg"+capitalize(key)] = itoaSlice(tuple[0]+BG_OFFSET, tuple[1]+BG_OFFSET)
		brightColors["bright"+capitalize(key)] = itoaSlice(tuple[0]+BRIGHT_OFFSET, tuple[1]+BRIGHT_OFFSET)
		brightBgColors["brightBg"+capitalize(key)] = itoaSlice(tuple[0]+BG_OFFSET+BRIGHT_OFFSET, tuple[1]+BG_OFFSET+BRIGHT_OFFSET)
	}
}

//	======
//	COLORS
//	======

//	FOREGROUND
//	----------

// Colors the string black
func Black(str string) string {
	return codes.New(colors["black"]).Wrap(str)
}

// Colors the string red
func Red(str string) string {
	return codes.New(colors["red"]).Wrap(str)
}

// Colors the string green
func Green(str string) string {
	return codes.New(colors["green"]).Wrap(str)
}

// Colors the string yellow
func Yellow(str string) string {
	return codes.New(colors["yellow"]).Wrap(str)
}

// Colors the string blue
func Blue(str string) string {
	return codes.New(colors["blue"]).Wrap(str)
}

// Colors the string magenta
func Magenta(str string) string {
	return codes.New(colors["magenta"]).Wrap(str)
}

// Colors the string cyan
func Cyan(str string) string {
	return codes.New(colors["cyan"]).Wrap(str)
}

// Colors the string white
func White(str string) string {
	return codes.New(colors["white"]).Wrap(str)
}

//	BRIGHT FOREGROUND
//	-----------------

// Colors the string brightBlack
func BrightBlack(str string) string {
	return codes.New(brightColors["brightBlack"]).Wrap(str)
}

// Colors the string brightRed
func BrightRed(str string) string {
	return codes.New(brightColors["brightRed"]).Wrap(str)
}

// Colors the string brightGreen
func BrightGreen(str string) string {
	return codes.New(brightColors["brightGreen"]).Wrap(str)
}

// Colors the string brightYellow
func BrightYellow(str string) string {
	return codes.New(brightColors["brightYellow"]).Wrap(str)
}

// Colors the string brightBlue
func BrightBlue(str string) string {
	return codes.New(brightColors["brightBlue"]).Wrap(str)
}

// Colors the string magenta
func BrightMagenta(str string) string {
	return codes.New(brightColors["brightMagenta"]).Wrap(str)
}

// Colors the string brightCyan
func BrightCyan(str string) string {
	return codes.New(brightColors["brightCyan"]).Wrap(str)
}

// Colors the string brightWhite
func BrightWhite(str string) string {
	return codes.New(brightColors["brightWhite"]).Wrap(str)
}

//	BACKGROUND
//	----------

// Colors the string bgBlack
func BgBlack(str string) string {
	return codes.New(bgColors["bgBlack"]).Wrap(str)
}

// Colors the string bgRed
func BgRed(str string) string {
	return codes.New(bgColors["bgRed"]).Wrap(str)
}

// Colors the string bgGreen
func BgGreen(str string) string {
	return codes.New(bgColors["bgGreen"]).Wrap(str)
}

// Colors the string bgYellow
func BgYellow(str string) string {
	return codes.New(bgColors["bgYellow"]).Wrap(str)
}

// Colors the string bgBlue
func BgBlue(str string) string {
	return codes.New(bgColors["bgBlue"]).Wrap(str)
}

// Colors the string magenta
func BgMagenta(str string) string {
	return codes.New(bgColors["bgMagenta"]).Wrap(str)
}

// Colors the string bgCyan
func BgCyan(str string) string {
	return codes.New(bgColors["bgCyan"]).Wrap(str)
}

// Colors the string bgWhite
func BgWhite(str string) string {
	return codes.New(bgColors["bgWhite"]).Wrap(str)
}

//	BRIGHT BACKGROUND
//	-----------------

// Colors the string brightBgBlack
func BrightBgBlack(str string) string {
	return codes.New(brightBgColors["brightBgBlack"]).Wrap(str)
}

// Colors the string brightBgRed
func BrightBgRed(str string) string {
	return codes.New(brightBgColors["brightBgRed"]).Wrap(str)
}

// Colors the string brightBgGreen
func BrightBgGreen(str string) string {
	return codes.New(brightBgColors["brightBgGreen"]).Wrap(str)
}

// Colors the string brightBgYellow
func BrightBgYellow(str string) string {
	return codes.New(brightBgColors["brightBgYellow"]).Wrap(str)
}

// Colors the string brightBgBlue
func BrightBgBlue(str string) string {
	return codes.New(brightBgColors["brightBgBlue"]).Wrap(str)
}

// Colors the string magenta
func BrightBgMagenta(str string) string {
	return codes.New(brightBgColors["brightBgMagenta"]).Wrap(str)
}

// Colors the string brightBgCyan
func BrightBgCyan(str string) string {
	return codes.New(brightBgColors["brightBgCyan"]).Wrap(str)
}

// Colors the string brightBgWhite
func BrightBgWhite(str string) string {
	return codes.New(brightBgColors["brightBgWhite"]).Wrap(str)
}

//	RGB
//	---

// Color the string with RGB values
func RGB(str string, r, g, b int) string {
	return codes.CSI + fmt.Sprintf("38;2;%d;%d;%dm%s%s", r, g, b, str, codes.RESET)
}

// Color the string bg with RGB values
func BgRGB(str string, r, g, b int) string {
	return codes.CSI + fmt.Sprintf("48;2;%d;%d;%dm%s%s", r, g, b, str, codes.RESET)
}
