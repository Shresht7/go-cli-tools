package main

import (
	"strings"
	"text/tabwriter"
)

type Table struct {
	data [][]string
}

func NewTable(data [][]string) *Table {
	return &Table{data}
}

func (t *Table) String() string {
	sb := strings.Builder{}
	tw := tabwriter.NewWriter(&sb, 0, 0, 1, ' ', tabwriter.AlignRight)

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

func main() {
	data := [][]string{
		{"Name", "Cost", "Quantity"},
		{"Apple", "10", "5"},
		{"Banana", "20", "10"},
		{"Orange", "15", "8"},
	}
	table := NewTable(data)
	println(table.String())
}
