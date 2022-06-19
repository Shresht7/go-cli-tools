package codes

//	-------------
//	MISCELLANEOUS
//	-------------

//	Link
func Link(text, url string) string {
	return OSC + "8;;" + url + BEL + text + OSC + "8;;" + BEL
}

//	Save Screen
func SaveScreen() string {
	return Escape("[?47h")
}

//	Load Screen
func LoadScreen() string {
	return Escape("[?47l")
}

//	Enable Alt Buffer
func EnableAltBuffer() string {
	return Escape("[?1049h")
}

//	Disable Alt Buffer
func DisableAltBuffer() string {
	return Escape("[?1049l")
}
