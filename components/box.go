package components

import (
	"strings"
)

type BoxBorder struct {
	top         string
	topRight    string
	right       string
	bottomRight string
	bottom      string
	bottomLeft  string
	left        string
	topLeft     string
}

// BorderSingle is a BoxBorder with single line characters
// ┌────────┐
// │        │
// └────────┘
var BorderSingle = BoxBorder{
	top:         "─",
	topRight:    "┐",
	right:       "│",
	bottomRight: "┘",
	bottom:      "─",
	bottomLeft:  "└",
	left:        "│",
	topLeft:     "┌",
}

type Box struct {
	content string
	border  BoxBorder
}

func NewBox(content string) *Box {
	return &Box{
		content: content,
		border:  BorderSingle,
	}
}

func (b *Box) SetBorder(border BoxBorder) {
	b.border = border
}

func (b *Box) String() string {

	// Split the content into lines
	lines := strings.Split(b.content, "\n")

	// Find the longest line
	maxWidth := getMaxWidth(lines)

	// Transform the content lines
	for i, line := range lines {
		lines[i] = b.border.left + " " + line + strings.Repeat(" ", maxWidth-len(line)) + " " + b.border.right
	}

	// Create the top border
	topBorder := b.border.topLeft + strings.Repeat(b.border.top, maxWidth+2) + b.border.topRight
	lines = append([]string{topBorder}, lines...)

	// Create the bottom border
	bottomBorder := b.border.bottomLeft + strings.Repeat(b.border.bottom, maxWidth+2) + b.border.bottomRight
	lines = append(lines, bottomBorder)

	// Join the lines
	return strings.Join(lines, "\n")

}

// getMaxWidth returns the length of the longest line in the given slice of strings
func getMaxWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}
