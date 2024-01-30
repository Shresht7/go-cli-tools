package main

import (
	"strings"
	"text/tabwriter"
)

// TABLE
// -----

type Table struct {
	headers []string
	data    [][]string

	minwidth int
	tabwidth int
	padding  int
	padchar  byte
	flags    uint
}

func NewTable(data [][]string) *Table {
	return &Table{
		data:    data,
		padding: 1,
		padchar: ' ',
	}
}

// CONFIGURATION
// -------------

func (t *Table) WithHeaders(headers []string) *Table {
	t.headers = headers
	return t
}

func (t *Table) WithMinWidth(minwidth int) *Table {
	t.minwidth = minwidth
	return t
}

func (t *Table) WithTabWidth(tabwidth int) *Table {
	t.tabwidth = tabwidth
	return t
}

func (t *Table) WithPadding(padding int) *Table {
	t.padding = padding
	return t
}

func (t *Table) WithPadChar(padchar byte) *Table {
	t.padchar = padchar
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

func (t *Table) String() string {
	sb := strings.Builder{}
	tw := tabwriter.NewWriter(&sb, t.minwidth, t.tabwidth, t.padding, t.padchar, t.flags)

	// Write headers
	for _, header := range t.headers {
		tw.Write([]byte(header))
		tw.Write([]byte("\t"))
	}
	tw.Write([]byte("\n"))

	// Write data
	for _, row := range t.data {
		for _, col := range row {
			tw.Write([]byte(col))
			tw.Write([]byte("\t"))
		}
		tw.Write([]byte("\n"))
	}

	tw.Flush()
	return sb.String()
}

// TODO: Remove this after testing
func main() {
	headers := []string{"Name", "Cost", "Quantity"}
	data := [][]string{
		{"Apple", "10", "5"},
		{"Banana", "20", "10"},
		{"Orange", "15", "8"},
	}

	table := NewTable(data).
		WithHeaders(headers).
		WithRightAlign(true)

	println(table.String())
}
