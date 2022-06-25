# go-cli-tools
--------------

<h2 align='center'>🚧Work In Progress🚧</h2>

Command-line tools and utilities for Go projects.

<details>

<summary>Table of Contents</summary>

- [# go-cli-tools](#-go-cli-tools)
- [📦 ANSI Modules](#-ansi-modules)
  - [🎨 Colors](#-colors)
  - [💄 Styles](#-styles)
  - [☝ Cursor](#-cursor)
  - [🧼 Clear](#-clear)
  - [RegExp](#regexp)
  - [Symbols](#symbols)
  - [Status](#status)
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

### 💄 Styles

The styles package allows you to format your text.

```go
import "github.com/Shresht7/go-cli-tools/ansi/styles"

fmt.Println(styles.Bold("GO"))
fmt.Println(styles.Hidden("Secret"))
```

### ☝ Cursor

The cursor package provides helper function for cursor manipulation.

```go
import "github.com/Shresht7/go-cli-tools/ansi/cursor"

fmt.Print(cursor.ToPos(15, 20))
fmt.Print(cursor.Up(3))
fmt.Print(cursor.ToNextLine(5))
```

### 🧼 Clear

The clear package provides helper functions to clear the terminal.

```go
import "github.com/Shresht7/go-cli-tools/ansi/clear"

fmt.Println(clear.EntireLine)
fmt.Println(clear.Screen)
```

### RegExp

Regular expressions to capture ansi codes. The `Strip` helper function can remove all ansi escape codes from a string.

```go
import "github.com/Shresht7/go-cli-tools/ansi/codes"

var text string = "..."
text = codes.Strip(text)
fmt.Println(text)
```

### Symbols

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

### Status

The `status` package includes some commonly used symbols with color.

<!-- TODO: Status Symbols Screenshot -->

```go
import "github.com/Shresht7/go-cli-tools/status"

fmt.Println(status.Tick, "Done")  //  ✔ Done
```

---

## 📑 License

> [MIT License](./LICENSE)
