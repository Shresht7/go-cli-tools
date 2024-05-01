package box_test

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/components/box"
)

func Example() {
	box := box.NewBox("Hello World!")
	fmt.Println(box)
	// Output:
	// ┌──────────────┐
	// │ Hello World! │
	// └──────────────┘
}
