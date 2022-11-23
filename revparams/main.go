package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	y := 0
	for index := range x {
		y = index
	}
	for i := y; i >= 1; i-- {
		for _, z := range x[i] {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}
