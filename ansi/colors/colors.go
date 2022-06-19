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
var colors = make(map[string][2]string)

//	Background Colors
var bgColors = make(map[string][2]string)

//	Bright Foreground Colors
var brightColors = make(map[string][2]string)

//	Bright Background Colors
var brightBgColors = make(map[string][2]string)

//	Formats given integers as a string tuple
func intsToStringTuple(x, y int) [2]string {
	return [2]string{strconv.Itoa(x), strconv.Itoa(y)}
}

//	Capitalizes a string
func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func init() {
	for key, tuple := range baseColors {
		colors[key] = intsToStringTuple(tuple[0], tuple[1])
		bgColors["bg"+capitalize(key)] = intsToStringTuple(tuple[0]+BG_OFFSET, tuple[1]+BG_OFFSET)
		brightColors["bright"+capitalize(key)] = intsToStringTuple(tuple[0]+BRIGHT_OFFSET, tuple[1]+BRIGHT_OFFSET)
		brightBgColors["brightBg"+capitalize(key)] = intsToStringTuple(tuple[0]+BG_OFFSET+BRIGHT_OFFSET, tuple[1]+BG_OFFSET+BRIGHT_OFFSET)
	}
}

//	======
//	COLORS
//	======

//	FOREGROUND
//	----------

//	Colors the string black
func Black(str string) string {
	return codes.Wrap(str, colors["black"])
}

//	Colors the string red
func Red(str string) string {
	return codes.Wrap(str, colors["red"])
}

//	Colors the string green
func Green(str string) string {
	return codes.Wrap(str, colors["green"])
}

//	Colors the string yellow
func Yellow(str string) string {
	return codes.Wrap(str, colors["yellow"])
}

//	Colors the string blue
func Blue(str string) string {
	return codes.Wrap(str, colors["blue"])
}

//	Colors the string magenta
func Magenta(str string) string {
	return codes.Wrap(str, colors["magenta"])
}

//	Colors the string cyan
func Cyan(str string) string {
	return codes.Wrap(str, colors["cyan"])
}

//	Colors the string white
func White(str string) string {
	return codes.Wrap(str, colors["white"])
}

//	BRIGHT FOREGROUND
//	-----------------

//	Colors the string brightBlack
func BrightBlack(str string) string {
	return codes.Wrap(str, colors["brightBlack"])
}

//	Colors the string brightRed
func BrightRed(str string) string {
	return codes.Wrap(str, colors["brightRed"])
}

//	Colors the string brightGreen
func BrightGreen(str string) string {
	return codes.Wrap(str, colors["brightGreen"])
}

//	Colors the string brightYellow
func BrightYellow(str string) string {
	return codes.Wrap(str, colors["brightYellow"])
}

//	Colors the string brightBlue
func BrightBlue(str string) string {
	return codes.Wrap(str, colors["brightBlue"])
}

//	Colors the string magenta
func BrightMagenta(str string) string {
	return codes.Wrap(str, colors["brightMagenta"])
}

//	Colors the string brightCyan
func BrightCyan(str string) string {
	return codes.Wrap(str, colors["brightCyan"])
}

//	Colors the string brightWhite
func BrightWhite(str string) string {
	return codes.Wrap(str, colors["brightWhite"])
}

//	BACKGROUND
//	----------

//	Colors the string bgBlack
func BgBlack(str string) string {
	return codes.Wrap(str, colors["bgBlack"])
}

//	Colors the string bgRed
func BgRed(str string) string {
	return codes.Wrap(str, colors["bgRed"])
}

//	Colors the string bgGreen
func BgGreen(str string) string {
	return codes.Wrap(str, colors["bgGreen"])
}

//	Colors the string bgYellow
func BgYellow(str string) string {
	return codes.Wrap(str, colors["bgYellow"])
}

//	Colors the string bgBlue
func BgBlue(str string) string {
	return codes.Wrap(str, colors["bgBlue"])
}

//	Colors the string magenta
func BgMagenta(str string) string {
	return codes.Wrap(str, colors["bgMagenta"])
}

//	Colors the string bgCyan
func BgCyan(str string) string {
	return codes.Wrap(str, colors["bgCyan"])
}

//	Colors the string bgWhite
func BgWhite(str string) string {
	return codes.Wrap(str, colors["bgWhite"])
}

//	BRIGHT BACKGROUND
//	-----------------

//	Colors the string brightBgBlack
func BrightBgBlack(str string) string {
	return codes.Wrap(str, colors["brightBgBlack"])
}

//	Colors the string brightBgRed
func BrightBgRed(str string) string {
	return codes.Wrap(str, colors["brightBgRed"])
}

//	Colors the string brightBgGreen
func BrightBgGreen(str string) string {
	return codes.Wrap(str, colors["brightBgGreen"])
}

//	Colors the string brightBgYellow
func BrightBgYellow(str string) string {
	return codes.Wrap(str, colors["brightBgYellow"])
}

//	Colors the string brightBgBlue
func BrightBgBlue(str string) string {
	return codes.Wrap(str, colors["brightBgBlue"])
}

//	Colors the string magenta
func BrightBgMagenta(str string) string {
	return codes.Wrap(str, colors["brightBgMagenta"])
}

//	Colors the string brightBgCyan
func BrightBgCyan(str string) string {
	return codes.Wrap(str, colors["brightBgCyan"])
}

//	Colors the string brightBgWhite
func BrightBgWhite(str string) string {
	return codes.Wrap(str, colors["brightBgWhite"])
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
