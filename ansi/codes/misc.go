package codes

//	-------------
//	MISCELLANEOUS
//	-------------

//	Link
func Link(text, url string) string {
	return OSC + "8;;" + url + BEL + text + OSC + "8;;" + BEL
}

const (
	//	Save the current screen
	SaveScreen = CSI + "?47h"
	//	Load the saved screen
	LoadScreen = CSI + "?47l"
	//	Enable alt buffer
	EnableAltBuffer = CSI + "?1049h"
	//	Disable alt buffer
	DisableAltBuffer = CSI + "?1049l"
)
