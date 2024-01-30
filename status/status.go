package status

import (
	"github.com/Shresht7/go-cli-tools/ansi/colors"
	"github.com/Shresht7/go-cli-tools/symbols"
)

var (
	Info     = colors.Blue(symbols.Info)          // ℹ
	Tick     = colors.Green(symbols.Tick)         // ✔
	Cross    = colors.Red(symbols.Cross)          // ✖
	Increase = colors.Green(symbols.TriangleUp)   // ▲
	Decrease = colors.Green(symbols.TriangleDown) // ▼
)
