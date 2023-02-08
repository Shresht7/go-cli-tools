package format

import (
	"fmt"
)

// HereDoc returns a here-document representation of the given string.
// See https://en.m.wikipedia.org/wiki/Here_document for more information.
func HereDoc(s string) string {
	return s
}

// HereDocf returns a here-document (https://en.m.wikipedia.org/wiki/Here_document)
// representation of the given string. The can be formatted using
// the given arguments as in fmt.Sprintf.
func HereDocf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}
