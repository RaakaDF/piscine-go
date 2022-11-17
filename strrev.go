package piscine

func StrRev(s string) string {
	Apia := []rune(s)
	for x, y := 0, len(Apia)-1; x < y; x, y = x+1, y-1 {
		Apia[x], Apia[y] = Apia[y], Apia[x]
	}
	return string(Apia)
}
