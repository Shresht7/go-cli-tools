# `go-cli-tools`

Command-line tools and utilities for Go projects.

[![Go Reference](https://pkg.go.dev/badge/github.com/Shresht7/go-cli-tools.svg)](https://pkg.go.dev/github.com/Shresht7/go-cli-tools)

<details>

<summary>Table of Contents</summary>

- [📦 ANSI Modules](#-ansi-modules)
  - [🎨 Colors](#-colors)
  - [💄 Styles](#-styles)
  - [☝ Cursor](#-cursor)
  - [🧼 Clear](#-clear)
  - [🅰 RegExp](#-regexp)
- [📚 Helpers](#-helpers)
  - [🏭 Composition](#-composition)
- [📃 Format](#-format)
  - [Align](#align)
  - [HereDoc](#heredoc)
  - [Pad](#pad)
- [📦 Layout](#-layout)
- [✔ Symbols](#-symbols)
- [✔ Status](#-status)
- [📑 License](#-license)

</details>

---

## 📦 ANSI Modules

### 🎨 Colors

The colors package allows you to style text with colors.

```go
import "github.com/Shresht7/go-cli-tools/ansi/colors"

fmt.Println(colors.Red("Giant"))
fmt.Println(colors.BgWhite("Dwarf"))
fmt.Println(colors.RGB("Supernova", 224, 65, 137))
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### 💄 Styles

The styles package allows you to format your text.

```go
import "github.com/Shresht7/go-cli-tools/ansi/styles"

fmt.Println(styles.Bold("GO"))
fmt.Println(styles.Hidden("Secret"))
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### ☝ Cursor

The cursor package provides helper function for cursor manipulation.

```go
import "github.com/Shresht7/go-cli-tools/ansi/cursor"

fmt.Print(cursor.ToPos(15, 20))
fmt.Print(cursor.Up(3))
fmt.Print(cursor.ToNextLine(5))
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### 🧼 Clear

The clear package provides helper functions to clear the terminal.

```go
import "github.com/Shresht7/go-cli-tools/ansi/clear"

fmt.Println(clear.EntireLine)
fmt.Println(clear.Screen)
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### 🅰 RegExp

Regular expressions to capture ansi codes. The `Strip` helper function can remove all ansi escape codes from a string.

```go
import "github.com/Shresht7/go-cli-tools/ansi/codes"

var text string = "..."
text = codes.Strip(text)
fmt.Println(text)
```

<div align="right">

  [⬆ Back to Top][top]

</div>

---

## 📚 Helpers

### 🏭 Composition

Composition helpers provide two utility functions compose and pipe that allow you to combine many ansi functions together.

```go
import "github.com/Shresht7/go-cli-tools/helpers"
import "github.com/Shresht7/go-cli-tools/ansi/colors"
import "github.com/Shresht7/go-cli-tools/ansi/styles"

styler = helpers.Compose(colors.Blue, styles.Bold, colors.BgWhite)
text = styler("Function Composition!")
```

<div align="right">

  [⬆ Back to Top][top]

</div>

---

## 📃 Format

### Align

Align text to the left, center or right.

```go
import "github.com/Shresht7/go-cli-tools/format"
var text string = "..."
fmt.Println(format.Align(text, &format.AlignOptions{ Mode: "Center" }))
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### HereDoc

HereDoc is a helper function to create a multi-line string.

```go
import "github.com/Shresht7/go-cli-tools/format"
var text string = format.HereDoc(`
    This is a multi-line string
    that can be used to create
      a template.
`)
// Output:
// This is a multi-line string
// that can be used to create
//  a template.
```

<div align="right">

  [⬆ Back to Top][top]

</div>

### Pad

Apply padding around the text.

```go
import "github.com/Shresht7/go-cli-tools/format"
var text string = "..."
fmt.Println(format.Pad(text, " ", 5))
```

<div align="right">

  [⬆ Back to Top][top]

</div>

---

## 📦 Layout

Create a table.

```go
import "github.com/Shresht7/go-cli-tools/layout"

var data [][]string // ...
table := layout.NewTable(data)
fmt.Println(table.String())
```

---

## ✔ Symbols

Unicode symbols for the terminal

```
✔ ℹ ⚠ ✖ ☰ ↑ ↓ ← → ♪ ♫ ■ ● ․ … › ▲ ▴ ▼ ▾ ◂ ▸ ⌂ ♥ ↔ ↕ ≈ ≠ ≤ ≥ ≡ ∞ ෴ ★ ▶ ⬢
```

```go
import "github.com/Shresht7/go-cli-tools/symbols"

fmt.Println(symbols.Warning, "Are you sure?")
//  ⚠ Are you sure?

fmt.Println(symbols.Tick, "Done")
//  ✔ Done

fmt.Println("Controls: ", symbols.ArrowUp, symbols.ArrowDown, symbol.ArrowLeft, symbols.ArrowRight)
//  Controls: ↑ ↓ ← →
```

<div align="right">

  [⬆ Back to Top][top]

</div>


## ✔ Status

The `status` package includes some commonly used symbols with color.

<!-- TODO: Status Symbols Screenshot -->

```go
import "github.com/Shresht7/go-cli-tools/status"

fmt.Println(status.Tick, "Done")  //  ✔ Done
```

<div align="right">

  [⬆ Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.

<!-- ====== -->
<!-- FOOTER -->
<!-- ====== -->

[top]: #go-cli-tools
