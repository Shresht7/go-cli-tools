package box

type BoxBorder struct {
	top         string
	topRight    string
	right       string
	bottomRight string
	bottom      string
	bottomLeft  string
	left        string
	topLeft     string
}

// BorderSingle is a BoxBorder with single line characters
// ┌────────┐
// │        │
// └────────┘
var BorderSingle = BoxBorder{
	top:         "─",
	topRight:    "┐",
	right:       "│",
	bottomRight: "┘",
	bottom:      "─",
	bottomLeft:  "└",
	left:        "│",
	topLeft:     "┌",
}

// BorderDouble is a BoxBorder with double line characters
// ╔════════╗
// ║        ║
// ╚════════╝
var BorderDouble = BoxBorder{
	top:         "═",
	topRight:    "╗",
	right:       "║",
	bottomRight: "╝",
	bottom:      "═",
	bottomLeft:  "╚",
	left:        "║",
	topLeft:     "╔",
}

// BorderRound is a BoxBorder with round characters
// ╭────────╮
// │        │
// ╰────────╯
var BorderRound = BoxBorder{
	top:         "─",
	topRight:    "╮",
	right:       "│",
	bottomRight: "╯",
	bottom:      "─",
	bottomLeft:  "╰",
	left:        "│",
	topLeft:     "╭",
}

// BorderBold is a BoxBorder with bold characters
// ┏━━━━━━━━┓
// ┃        ┃
// ┗━━━━━━━━┛
var BorderBold = BoxBorder{
	top:         "━",
	topRight:    "┓",
	right:       "┃",
	bottomRight: "┛",
	bottom:      "━",
	bottomLeft:  "┗",
	left:        "┃",
	topLeft:     "┏",
}

// BorderClassic is a BoxBorder with classic characters
// +--------+
// |        |
// +--------+
var BorderClassic = BoxBorder{
	top:         "-",
	topRight:    "+",
	right:       "|",
	bottomRight: "+",
	bottom:      "-",
	bottomLeft:  "+",
	left:        "|",
	topLeft:     "+",
}
