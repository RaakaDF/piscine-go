package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, y := range s {
		z01.PrintRune(y)
	}
}
