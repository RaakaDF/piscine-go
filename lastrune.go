package piscine

func LastRune(s string) rune {
	x := []rune(s)
	y := len([]rune(s))
	return x[y-1]
}
