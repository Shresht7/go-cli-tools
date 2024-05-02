package frames

type Spinner struct {
	Interval int
	Frames   []string
}

// Map applies a function to each frame of the spinner.
func (s *Spinner) Map(f func(int, string) string) *Spinner {
	for i, frame := range s.Frames {
		s.Frames[i] = f(i, frame)
	}
	return s
}

var Windows = Spinner{
	Interval: 80,
	Frames: []string{
		"/",
		"-",
		"\\",
		"|",
	},
}

var Dots = Spinner{
	Interval: 80,
	Frames: []string{
		"⠋",
		"⠙",
		"⠹",
		"⠸",
		"⠼",
		"⠴",
		"⠦",
		"⠧",
		"⠇",
		"⠏",
	},
}

var Dots2 = Spinner{
	Interval: 80,
	Frames: []string{
		"⣾",
		"⣽",
		"⣻",
		"⢿",
		"⡿",
		"⣟",
		"⣯",
		"⣷",
	},
}

var Dots3 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠋",
		"⠙",
		"⠚",
		"⠞",
		"⠖",
		"⠦",
		"⠴",
		"⠲",
		"⠳",
		"⠓",
	},
}

var Dots4 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠄",
		"⠆",
		"⠇",
		"⠋",
		"⠙",
		"⠸",
		"⠰",
		"⠠",
		"⠰",
		"⠸",
		"⠙",
		"⠋",
		"⠇",
		"⠆",
	},
}

var Dots5 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠋",
		"⠙",
		"⠚",
		"⠒",
		"⠂",
		"⠂",
		"⠒",
		"⠲",
		"⠴",
		"⠦",
		"⠖",
		"⠒",
		"⠐",
		"⠐",
		"⠒",
		"⠓",
		"⠋",
	},
}

var Dots6 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠁",
		"⠉",
		"⠙",
		"⠚",
		"⠒",
		"⠂",
		"⠂",
		"⠒",
		"⠲",
		"⠴",
		"⠤",
		"⠄",
		"⠄",
		"⠤",
		"⠴",
		"⠲",
		"⠒",
		"⠂",
		"⠂",
		"⠒",
		"⠚",
		"⠙",
		"⠉",
		"⠁",
	},
}

var Dots7 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠈",
		"⠉",
		"⠋",
		"⠓",
		"⠒",
		"⠐",
		"⠐",
		"⠒",
		"⠖",
		"⠦",
		"⠤",
		"⠠",
		"⠠",
		"⠤",
		"⠦",
		"⠖",
		"⠒",
		"⠐",
		"⠐",
		"⠒",
		"⠓",
		"⠋",
		"⠉",
		"⠈",
	},
}

var Dots8 = Spinner{
	Interval: 80,
	Frames: []string{
		"⠁",
		"⠁",
		"⠉",
		"⠙",
		"⠚",
		"⠒",
		"⠂",
		"⠂",
		"⠒",
		"⠲",
		"⠴",
		"⠤",
		"⠄",
		"⠄",
		"⠤",
		"⠠",
		"⠠",
		"⠤",
		"⠦",
		"⠖",
		"⠒",
		"⠐",
		"⠐",
		"⠒",
		"⠓",
		"⠋",
		"⠉",
		"⠈",
		"⠈",
	},
}

var Dots9 = Spinner{
	Interval: 80,
	Frames: []string{
		"⢹",
		"⢺",
		"⢼",
		"⣸",
		"⣇",
		"⡧",
		"⡗",
		"⡏",
	},
}

var Dots10 = Spinner{
	Interval: 80,
	Frames: []string{
		"⢄",
		"⢂",
		"⢁",
		"⡁",
		"⡈",
		"⡐",
		"⡠",
	},
}

var Dots11 = Spinner{
	Interval: 100,
	Frames: []string{
		"⠁",
		"⠂",
		"⠄",
		"⡀",
		"⢀",
		"⠠",
		"⠐",
		"⠈",
	},
}

var Dots12 = Spinner{
	Interval: 80,
	Frames: []string{
		"⢀⠀",
		"⡀⠀",
		"⠄⠀",
		"⢂⠀",
		"⡂⠀",
		"⠅⠀",
		"⢃⠀",
		"⡃⠀",
		"⠍⠀",
		"⢋⠀",
		"⡋⠀",
		"⠍⠁",
		"⢋⠁",
		"⡋⠁",
		"⠍⠉",
		"⠋⠉",
		"⠋⠉",
		"⠉⠙",
		"⠉⠙",
		"⠉⠩",
		"⠈⢙",
		"⠈⡙",
		"⢈⠩",
		"⡀⢙",
		"⠄⡙",
		"⢂⠩",
		"⡂⢘",
		"⠅⡘",
		"⢃⠨",
		"⡃⢐",
		"⠍⡐",
		"⢋⠠",
		"⡋⢀",
		"⠍⡁",
		"⢋⠁",
		"⡋⠁",
		"⠍⠉",
		"⠋⠉",
		"⠋⠉",
		"⠉⠙",
		"⠉⠙",
		"⠉⠩",
		"⠈⢙",
		"⠈⡙",
		"⠈⠩",
		"⠀⢙",
		"⠀⡙",
		"⠀⠩",
		"⠀⢘",
		"⠀⡘",
		"⠀⠨",
		"⠀⢐",
		"⠀⡐",
		"⠀⠠",
		"⠀⢀",
		"⠀⡀",
	},
}

var Line = Spinner{
	Interval: 130,
	Frames: []string{
		"─",
		"\\",
		"|",
		"/",
	},
}

var Line2 = Spinner{
	Interval: 100,
	Frames: []string{
		"⠂",
		"-",
		"–",
		"—",
		"–",
		"-",
	},
}

var Pipe = Spinner{
	Interval: 100,
	Frames: []string{
		"┤",
		"┘",
		"┴",
		"└",
		"├",
		"┌",
		"┬",
		"┐",
	},
}

var SimpleDots = Spinner{
	Interval: 400,
	Frames: []string{
		".  ",
		".. ",
		"...",
		"   ",
	},
}

var SimpleDotsScrolling = Spinner{
	Interval: 200,
	Frames: []string{
		".  ",
		".. ",
		"...",
		" ..",
		"  .",
	},
}

var Star = Spinner{
	Interval: 70,
	Frames: []string{
		"✶",
		"✸",
		"✹",
		"✺",
		"✹",
		"✷",
	},
}

var Star2 = Spinner{
	Interval: 80,
	Frames: []string{
		"+",
		"x",
		"*",
	},
}

var Flip = Spinner{
	Interval: 70,
	Frames: []string{
		"_",
		"_",
		"_",
		"-",
		"`",
		"`",
		"'",
		"´",
		"-",
		"_",
		"_",
		"_",
	},
}

var Hamburger = Spinner{
	Interval: 100,
	Frames: []string{
		"☱",
		"☲",
		"☴",
	},
}

var GrowVertical = Spinner{
	Interval: 120,
	Frames: []string{
		"▁",
		"▃",
		"▄",
		"▅",
		"▆",
		"▇",
		"▆",
		"▅",
		"▄",
		"▃",
	},
}

var GrowHorizontal = Spinner{
	Interval: 120,
	Frames: []string{
		"▏",
		"▎",
		"▍",
		"▌",
		"▋",
		"▊",
		"▉",
		"▊",
		"▋",
		"▌",
		"▍",
		"▎",
	},
}

var Noise = Spinner{
	Interval: 100,
	Frames: []string{
		"▓",
		"▒",
		"░",
	},
}

var Bouncing = Spinner{
	Interval: 120,
	Frames: []string{
		"⠁",
		"⠂",
		"⠄",
		"⠂",
	},
}

var BoxBounce = Spinner{
	Interval: 120,
	Frames: []string{
		"▖",
		"▘",
		"▝",
		"▗",
	},
}

var BoxBounce2 = Spinner{
	Interval: 100,
	Frames: []string{
		"▌",
		"▀",
		"▐",
		"▄",
	},
}

var Triangle = Spinner{
	Interval: 50,
	Frames: []string{
		"◢",
		"◣",
		"◤",
		"◥",
	},
}

var Arc = Spinner{
	Interval: 100,
	Frames: []string{
		"◜",
		"◠",
		"◝",
		"◞",
		"◡",
		"◟",
	},
}

var Circle = Spinner{
	Interval: 120,
	Frames: []string{
		"◡",
		"⊙",
		"◠",
	},
}

var SquareCorners = Spinner{
	Interval: 180,
	Frames: []string{
		"◰",
		"◳",
		"◲",
		"◱",
	},
}

var CircleQuarters = Spinner{
	Interval: 120,
	Frames: []string{
		"◴",
		"◷",
		"◶",
		"◵",
	},
}

var CircleHalves = Spinner{
	Interval: 50,
	Frames: []string{
		"◐",
		"◓",
		"◑",
		"◒",
	},
}

var Arrows = Spinner{
	Interval: 100,
	Frames: []string{
		"←",
		"↖",
		"↑",
		"↗",
		"→",
		"↘",
		"↓",
		"↙",
	},
}

var Arrows2 = Spinner{
	Interval: 80,
	Frames: []string{
		"⬆️ ",
		"↗️ ",
		"➡️ ",
		"↘️ ",
		"⬇️ ",
		"↙️ ",
		"⬅️ ",
		"↖️ ",
	},
}

var Arrows3 = Spinner{
	Interval: 120,
	Frames: []string{
		"▹▹▹▹▹",
		"▸▹▹▹▹",
		"▹▸▹▹▹",
		"▹▹▸▹▹",
		"▹▹▹▸▹",
		"▹▹▹▹▸",
	},
}

var BouncingBar = Spinner{
	Interval: 80,
	Frames: []string{
		"[    ]",
		"[=   ]",
		"[==  ]",
		"[=== ]",
		"[ ===]",
		"[  ==]",
		"[   =]",
		"[    ]",
		"[   =]",
		"[  ==]",
		"[ ===]",
		"[====]",
		"[=== ]",
		"[==  ]",
		"[=   ]",
	},
}

var BouncingBall = Spinner{
	Interval: 80,
	Frames: []string{
		"( ●    )",
		"(  ●   )",
		"(   ●  )",
		"(    ● )",
		"(     ●)",
		"(    ● )",
		"(   ●  )",
		"(  ●   )",
		"( ●    )",
		"(●     )",
	},
}

var Smiley = Spinner{
	Interval: 200,
	Frames: []string{
		"😄 ",
		"😝 ",
	},
}

var Hearts = Spinner{
	Interval: 100,
	Frames: []string{
		"💛 ",
		"💙 ",
		"💜 ",
		"💚 ",
		"❤️ ",
	},
}

var Clock = Spinner{
	Interval: 100,
	Frames: []string{
		"🕛 ",
		"🕐 ",
		"🕑 ",
		"🕒 ",
		"🕓 ",
		"🕔 ",
		"🕕 ",
		"🕖 ",
		"🕗 ",
		"🕘 ",
		"🕙 ",
		"🕚 ",
	},
}

var Earth = Spinner{
	Interval: 180,
	Frames: []string{
		"🌍 ",
		"🌎 ",
		"🌏 ",
	},
}

var Moon = Spinner{
	Interval: 80,
	Frames: []string{
		"🌑 ",
		"🌒 ",
		"🌓 ",
		"🌔 ",
		"🌕 ",
		"🌖 ",
		"🌗 ",
		"🌘 ",
	},
}

var Runner = Spinner{
	Interval: 140,
	Frames: []string{
		"🚶 ",
		"🏃 ",
	},
}

var Pong = Spinner{
	Interval: 80,
	Frames: []string{
		"▐⠂       ▌",
		"▐⠈       ▌",
		"▐ ⠂      ▌",
		"▐ ⠠      ▌",
		"▐  ⡀     ▌",
		"▐  ⠠     ▌",
		"▐   ⠂    ▌",
		"▐   ⠈    ▌",
		"▐    ⠂   ▌",
		"▐    ⠠   ▌",
		"▐     ⡀  ▌",
		"▐     ⠠  ▌",
		"▐      ⠂ ▌",
		"▐      ⠈ ▌",
		"▐       ⠂▌",
		"▐       ⠠▌",
		"▐       ⡀▌",
		"▐      ⠠ ▌",
		"▐      ⠂ ▌",
		"▐     ⠈  ▌",
		"▐     ⠂  ▌",
		"▐    ⠠   ▌",
		"▐    ⡀   ▌",
		"▐   ⠠    ▌",
		"▐   ⠂    ▌",
		"▐  ⠈     ▌",
		"▐  ⠂     ▌",
		"▐ ⠠      ▌",
		"▐ ⡀      ▌",
		"▐⠠       ▌",
	},
}

var Shark = Spinner{
	Interval: 120,
	Frames: []string{
		"▐|\\____________▌",
		"▐_|\\___________▌",
		"▐__|\\__________▌",
		"▐___|\\_________▌",
		"▐____|\\________▌",
		"▐_____|\\_______▌",
		"▐______|\\______▌",
		"▐_______|\\_____▌",
		"▐________|\\____▌",
		"▐_________|\\___▌",
		"▐__________|\\__▌",
		"▐___________|\\_▌",
		"▐____________|\\▌",
		"▐____________/|▌",
		"▐___________/|_▌",
		"▐__________/|__▌",
		"▐_________/|___▌",
		"▐________/|____▌",
		"▐_______/|_____▌",
		"▐______/|______▌",
		"▐_____/|_______▌",
		"▐____/|________▌",
		"▐___/|_________▌",
		"▐__/|__________▌",
		"▐_/|___________▌",
		"▐/|____________▌",
	},
}

var Weather = Spinner{
	Interval: 100,
	Frames: []string{
		"☀️ ",
		"☀️ ",
		"☀️ ",
		"🌤 ",
		"⛅️ ",
		"🌥 ",
		"☁️ ",
		"🌧 ",
		"🌨 ",
		"🌧 ",
		"🌨 ",
		"🌧 ",
		"🌨 ",
		"⛈ ",
		"🌨 ",
		"🌧 ",
		"🌨 ",
		"☁️ ",
		"🌥 ",
		"⛅️ ",
		"🌤 ",
		"☀️ ",
		"☀️ ",
	},
}

var Point = Spinner{
	Interval: 125,
	Frames: []string{
		"∙∙∙",
		"●∙∙",
		"∙●∙",
		"∙∙●",
		"∙∙∙",
	},
}

var Layer = Spinner{
	Interval: 150,
	Frames: []string{
		"-",
		"=",
		"≡",
	},
}
