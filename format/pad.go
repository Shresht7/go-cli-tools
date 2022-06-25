package format

import "strings"

//	=======
//	PADDING
//	=======

//	Apply padding around the string
func Pad(s, char string, count int) string {
	space := count / 2
	if space%2 == 0 {
		return strings.Repeat(char, space) + s + strings.Repeat(char, space)
	} else {
		return strings.Repeat(char, space) + s + strings.Repeat(char, space+1)
	}
}

//	Apply padding to the left of the string
func PadLeft(s, char string, count int) string {
	return strings.Repeat(char, count) + s
}

//	Apply padding to the right of the string
func PadRight(s, char string, count int) string {
	return s + strings.Repeat(char, count)
}
