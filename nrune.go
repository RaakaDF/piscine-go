package main

func NRune(s string, n int) rune {
	x := []rune(s)
	y := len([]rune(s))
	if (n <= 0) || (n > y) {
		return 0
	} else {
		return x[n-1]
	}
}

//func main() {
//	z01.PrintRune(NRune("Hello!", 3))
//	z01.PrintRune(NRune("Salut!", 2))
//	z01.PrintRune(NRune("Bye!", -1))
//	z01.PrintRune(NRune("Bye!", 5))
//	z01.PrintRune(NRune("Ola!", 4))
//	z01.PrintRune('\n')
//}
