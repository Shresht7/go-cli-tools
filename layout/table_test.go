package layout_test

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/layout"
)

func Example() {
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "25", "San Francisco"},
		{"Bob", "30", "Los Angeles"},
		{"Charlie", "35", "New York"},
	}

	table := layout.NewTable(data).WithPadChar('_')

	// Note: We use _ instead of space for padding because the formatter trims trailing spaces
	// in the Output block automatically, making it hard to test.

	fmt.Println(table)
	// Output:
	// Name____Age_City__________
	// Alice___25__San Francisco_
	// Bob_____30__Los Angeles___
	// Charlie_35__New York______
}

func ExampleTable() {
	data := [][]string{
		{"Alice", "25", "San Francisco"},
		{"Bob", "30", "Los Angeles"},
		{"Charlie", "35", "New York"},
	}

	table := layout.NewTable(data).
		WithHeaders([]string{"Name", "Age", "City"}).
		WithMinWidth(20).
		WithTabWidth(4).
		WithPadding(3).
		WithPadChar('.')

	fmt.Println(table)
	// Output:
	// Name................Age.................City................
	// Alice...............25..................San Francisco.......
	// Bob.................30..................Los Angeles.........
	// Charlie.............35..................New York............
}
