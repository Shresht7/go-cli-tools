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
