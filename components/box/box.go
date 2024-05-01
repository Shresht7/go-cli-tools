package box

import (
	"strings"

	"github.com/Shresht7/go-cli-tools/format"
	"github.com/Shresht7/go-cli-tools/helpers"
)

// Box represents a box with content and a border
type Box struct {
	content string // The content of the box

	title          string // The title of the box
	titleAlignment string // The alignment of the title

	border BoxBorder // The borders of the box
}

// Instantiates a new Box with the given content
func New(content string) *Box {
	return &Box{
		title:   "",
		content: content,
		border:  BorderSingle,
	}
}

// Sets the content of the box
func (b *Box) WithContent(content string) *Box {
	b.content = content
	return b
}

// Sets the title of the box
func (b *Box) WithTitle(title string) *Box {
	b.title = title
	return b
}

// Sets the alignment of the title
func (b *Box) WithTitleAlignment(alignment string) *Box {
	b.titleAlignment = alignment
	return b
}

// Sets the borders of the box
func (b *Box) WithBorder(border BoxBorder) *Box {
	b.border = border
	return b
}

// Returns the string representation of the box
func (b *Box) String() string {

	// Split the content into lines
	lines := strings.Split(b.content, "\n")

	// Find the longest line
	maxWidth := helpers.GetMaxLineWidth(lines)

	// Adjust the content lines
	lines = b.adjustContent(lines, maxWidth)

	// Create the top border
	topBorder := b.createTopBorder(maxWidth)
	lines = append([]string{topBorder}, lines...)

	// Create the bottom border
	bottomBorder := b.createBottomBorder(maxWidth)
	lines = append(lines, bottomBorder)

	// Join the lines
	return strings.Join(lines, "\n")

}

// Creates the top border of the box
func (b *Box) createTopBorder(maxWidth int) string {
	if b.title != "" {
		title := format.Align(" "+b.title+" ", &format.AlignOptions{
			Mode:  b.titleAlignment,
			Pad:   b.border.top,
			Width: maxWidth,
		})
		return b.border.topLeft + b.border.top + title + b.border.top + b.border.topRight
	}
	return b.border.topLeft + strings.Repeat(b.border.top, maxWidth+2) + b.border.topRight
}

// Adjusts the content lines for the box
func (b *Box) adjustContent(lines []string, maxWidth int) []string {
	// Transform the content lines
	for i, line := range lines {
		lines[i] = b.border.left + " " + line + strings.Repeat(" ", maxWidth-len(line)) + " " + b.border.right
	}
	return lines
}

// Creates the bottom border of the box
func (b *Box) createBottomBorder(maxWidth int) string {
	return b.border.bottomLeft + strings.Repeat(b.border.bottom, maxWidth+2) + b.border.bottomRight
}
