package layout

import (
	"strings"
	"text/tabwriter"
)

// TABLE
// -----

type Table struct {
	headers []string
	data    [][]string

	minWidth int
	tabWidth int
	padding  int
	padChar  byte
	flags    uint
}

func NewTable(data [][]string) *Table {
	return &Table{
		data:    data,
		padding: 1,
		padChar: ' ',
	}
}

// CONFIGURATION
// -------------

func (t *Table) WithHeaders(headers []string) *Table {
	t.headers = headers
	return t
}

func (t *Table) WithMinWidth(minWidth int) *Table {
	t.minWidth = minWidth
	return t
}

func (t *Table) WithTabWidth(tabWidth int) *Table {
	t.tabWidth = tabWidth
	return t
}

func (t *Table) WithPadding(padding int) *Table {
	t.padding = padding
	return t
}

func (t *Table) WithPadChar(padChar byte) *Table {
	t.padChar = padChar
	return t
}

func (t *Table) WithRightAlign(b bool) *Table {
	if b {
		t.flags = tabwriter.AlignRight
	}
	return t
}

// RENDER
// ------

func writeRow(tw *tabwriter.Writer, row []string) {
	for _, col := range row {
		tw.Write([]byte(col))
		tw.Write([]byte("\t"))
	}
	tw.Write([]byte("\n"))
}

func (t *Table) String() string {
	sb := strings.Builder{}
	tw := tabwriter.NewWriter(&sb, t.minWidth, t.tabWidth, t.padding, t.padChar, t.flags)

	// Write headers
	if len(t.headers) > 0 {
		writeRow(tw, t.headers)
	}

	// Write data
	for _, row := range t.data {
		writeRow(tw, row)
	}

	tw.Flush()
	return sb.String()
}
