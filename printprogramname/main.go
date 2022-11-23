package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	for index, y := range x[0] {
		if index >= 0 && y != '/' && y != '.' {
			i := y
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
