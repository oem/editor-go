package main

import (
	"fmt"

	"github.com/oem/replay/piecetable"
)

func main() {
	fmt.Println("replay: a minimal modal editor")
	pt := piecetable.New("an example string")
	pt.Insert("unremarkable ", 3)
	fmt.Println(pt, pt.Get())
}
