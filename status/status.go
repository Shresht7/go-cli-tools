package status

import (
	"github.com/Shresht7/go-cli-tools/ansi/colors"
	"github.com/Shresht7/go-cli-tools/symbols"
)

var (
	// ℹ
	Info = colors.Blue(symbols.Info)

	// ✔
	Tick    = colors.Green(symbols.Tick)
	Done    = colors.Green(symbols.Tick)
	Success = colors.Green(symbols.Tick)

	// ✖
	Error  = colors.Red(symbols.Cross)
	Close  = colors.Red(symbols.Cross)
	Cancel = colors.Red(symbols.Cross)
	Cross  = colors.Red(symbols.Cross)

	// ▲
	TriangleUp = colors.Green(symbols.TriangleUp)
	Increase   = colors.Green(symbols.TriangleUp)

	// ▼
	TriangleDown = colors.Red(symbols.TriangleDown)
	Decrease     = colors.Green(symbols.TriangleDown)
)
