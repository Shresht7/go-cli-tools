package box_test

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/components/box"
)

func Example() {
	box := box.New("Hello World!")
	fmt.Println(box)
	// Output:
	// ┌──────────────┐
	// │ Hello World! │
	// └──────────────┘
}

func Title_Example() {
	box := box.New("Hello World!")
	box.SetTitle("Title")
	box.SetTitleAlignment("Left")
	fmt.Println(box)
	// Output:
	// ┌─ Title ──────┐
	// │ Hello World! │
	// └──────────────┘
}
