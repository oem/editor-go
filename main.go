package main

import (
	"fmt"

	"github.com/oem/replay/piecetable"
)

func main() {
	fmt.Println("editor-go: tools for text editing")
	pt := piecetable.New("an example string")
	pt.Insert("unremarkable ", 3)
	fmt.Println(pt, pt.Get())
}
