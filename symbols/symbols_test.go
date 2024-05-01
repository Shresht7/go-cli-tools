package symbols_test

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/symbols"
)

func Example() {
	fmt.Println(symbols.Info)
	fmt.Println(symbols.Tick)
	fmt.Println(symbols.Cross)
	fmt.Println(symbols.TriangleUp)
	fmt.Println(symbols.TriangleDown)
	fmt.Println(symbols.Warning)
	// Output:
	// ℹ
	// ✔
	// ✖
	// ▲
	// ▼
	// ⚠
}
