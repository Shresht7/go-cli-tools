package components_test

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/components"
)

func Example() {
	box := components.NewBox("Hello World!")
	fmt.Println(box)
	// Output:
	// ┌──────────────┐
	// │ Hello World! │
	// └──────────────┘
}
