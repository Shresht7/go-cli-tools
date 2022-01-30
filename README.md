# go-cli-tools
--------------

Command-line tools and utilities.

## Library

### Colors

The colors package allows you to style text with colors.

```go
import "github.com/Shresht7/go-cli-tools/colors"

fmt.Println(colors.Red("Hello"))
```

### Styles

The styles package allows you to format your text.

```go
import "github.com/Shresht7/go-cli-tools/styles"

fmt.Println(styles.Bold("GO"))
```

### Cursor

The cursor package provides helper function for cursor manipulation

```Go
import "github.com/Shresht7/go-cli-tools/cursor"

fmt.Print(cursor.ToPos(15, 20))
```