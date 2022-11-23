package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args[0]
	for _, y := range x {
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}
