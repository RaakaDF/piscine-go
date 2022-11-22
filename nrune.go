package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	y := len([]rune(s))
	if (n <= 0) || (n > y) {
		return 0
	} else {
		return x[n-1]
	}
}
