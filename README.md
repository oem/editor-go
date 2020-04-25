![replay](docs/logo.png)

# Replay

This package implements data structures commonly used for text editing.

## Piece Table

The piece table is the first such data structure, have a look at the piecetable package for more details.

Example usage:

```go
import "github.com/oem/replay/piecetable"

pt := piecetable.New("an example string")
pt.Insert("unremarkable ", 3)
fmt.Println(pt, pt.Get()) 
// &{an example string unremarkable  [{0 0 false} {0 3 false} {0 13 true} {3 14 false}]} an unremarkable example string
```
