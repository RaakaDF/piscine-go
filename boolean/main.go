package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) bool {
	return nbr%2 == 0
}

func isEven(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		EvenMsg := "I have an even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"

		printStr(OddMsg)

	}
}
