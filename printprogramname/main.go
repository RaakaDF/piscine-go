package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args[0]
	for index, word := range x {
		if index >= 0 && word != '/' && word != '.' {
			i := word
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
