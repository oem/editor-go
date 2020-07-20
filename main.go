package main

import (
	"fmt"

	"github.com/oem/editor-go/piecetable"
)

func main() {
	fmt.Println("editor-go: tools for text editing")
	pt := piecetable.New("an example string")
	_ = pt.Insert("unremarkable ", 3)
	fmt.Println(pt, pt.Get())
}
