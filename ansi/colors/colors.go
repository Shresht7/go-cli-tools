package colors

import (
	"strconv"
	"strings"

	"github.com/Shresht7/go-cli-tools/ansi/codes"
)

//	Base Colors
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

//	Background  color offset
const BG_OFFSET = 10

//	Bright color offset
const BRIGHT_OFFSET = 60

//	Foreground Colors
var colors = make(map[string][]string)

//	Background Colors
var bgColors = make(map[string][]string)

//	Bright Foreground Colors
var brightColors = make(map[string][]string)

//	Bright Background Colors
var brightBgColors = make(map[string][]string)

//	Formats given integers as a string tuple
func itoaSlice(i ...int) []string {
	var s []string
	for _, v := range i {
		s = append(s, strconv.Itoa(v))
	}
	return s
}

//	Capitalizes a string
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

//	Colors the string black
func Black(str string) string {
	return codes.Wrap(str, colors["black"], true)
}

//	Colors the string red
func Red(str string) string {
	return codes.Wrap(str, colors["red"], true)
}

//	Colors the string green
func Green(str string) string {
	return codes.Wrap(str, colors["green"], true)
}

//	Colors the string yellow
func Yellow(str string) string {
	return codes.Wrap(str, colors["yellow"], true)
}

//	Colors the string blue
func Blue(str string) string {
	return codes.Wrap(str, colors["blue"], true)
}

//	Colors the string magenta
func Magenta(str string) string {
	return codes.Wrap(str, colors["magenta"], true)
}

//	Colors the string cyan
func Cyan(str string) string {
	return codes.Wrap(str, colors["cyan"], true)
}

//	Colors the string white
func White(str string) string {
	return codes.Wrap(str, colors["white"], true)
}

//	BRIGHT FOREGROUND
//	-----------------

//	Colors the string brightBlack
func BrightBlack(str string) string {
	return codes.Wrap(str, colors["brightBlack"], true)
}

//	Colors the string brightRed
func BrightRed(str string) string {
	return codes.Wrap(str, colors["brightRed"], true)
}

//	Colors the string brightGreen
func BrightGreen(str string) string {
	return codes.Wrap(str, colors["brightGreen"], true)
}

//	Colors the string brightYellow
func BrightYellow(str string) string {
	return codes.Wrap(str, colors["brightYellow"], true)
}

//	Colors the string brightBlue
func BrightBlue(str string) string {
	return codes.Wrap(str, colors["brightBlue"], true)
}

//	Colors the string magenta
func BrightMagenta(str string) string {
	return codes.Wrap(str, colors["brightMagenta"], true)
}

//	Colors the string brightCyan
func BrightCyan(str string) string {
	return codes.Wrap(str, colors["brightCyan"], true)
}

//	Colors the string brightWhite
func BrightWhite(str string) string {
	return codes.Wrap(str, colors["brightWhite"], true)
}

//	BACKGROUND
//	----------

//	Colors the string bgBlack
func BgBlack(str string) string {
	return codes.Wrap(str, colors["bgBlack"], true)
}

//	Colors the string bgRed
func BgRed(str string) string {
	return codes.Wrap(str, colors["bgRed"], true)
}

//	Colors the string bgGreen
func BgGreen(str string) string {
	return codes.Wrap(str, colors["bgGreen"], true)
}

//	Colors the string bgYellow
func BgYellow(str string) string {
	return codes.Wrap(str, colors["bgYellow"], true)
}

//	Colors the string bgBlue
func BgBlue(str string) string {
	return codes.Wrap(str, colors["bgBlue"], true)
}

//	Colors the string magenta
func BgMagenta(str string) string {
	return codes.Wrap(str, colors["bgMagenta"], true)
}

//	Colors the string bgCyan
func BgCyan(str string) string {
	return codes.Wrap(str, colors["bgCyan"], true)
}

//	Colors the string bgWhite
func BgWhite(str string) string {
	return codes.Wrap(str, colors["bgWhite"], true)
}

//	BRIGHT BACKGROUND
//	-----------------

//	Colors the string brightBgBlack
func BrightBgBlack(str string) string {
	return codes.Wrap(str, colors["brightBgBlack"], true)
}

//	Colors the string brightBgRed
func BrightBgRed(str string) string {
	return codes.Wrap(str, colors["brightBgRed"], true)
}

//	Colors the string brightBgGreen
func BrightBgGreen(str string) string {
	return codes.Wrap(str, colors["brightBgGreen"], true)
}

//	Colors the string brightBgYellow
func BrightBgYellow(str string) string {
	return codes.Wrap(str, colors["brightBgYellow"], true)
}

//	Colors the string brightBgBlue
func BrightBgBlue(str string) string {
	return codes.Wrap(str, colors["brightBgBlue"], true)
}

//	Colors the string magenta
func BrightBgMagenta(str string) string {
	return codes.Wrap(str, colors["brightBgMagenta"], true)
}

//	Colors the string brightBgCyan
func BrightBgCyan(str string) string {
	return codes.Wrap(str, colors["brightBgCyan"], true)
}

//	Colors the string brightBgWhite
func BrightBgWhite(str string) string {
	return codes.Wrap(str, colors["brightBgWhite"], true)
}

//	RGB
//	---

//	Color the string with RGB values
func RGB(str string, r, g, b int) string {
	return codes.Escape("[38;2;%d;%d;%dm%s%s", r, g, b, str, codes.RESET)
}

//	Color the string bg with RGB values
func BgRGB(str string, r, g, b int) string {
	return codes.Escape("[48;2;%d;%d;%dm%s%s", r, g, b, str, codes.RESET)
}
