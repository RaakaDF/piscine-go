package piscine

func SplitWhiteSpaces(n string) []string {
	var s []string
	x := 1
	re := 0
	result := ""
	for c := range n {
		if isWhiteSpace(n[c]) {
			x++
		}
		re++
	}
	s = make([]string, x)
	i := 0
	for j, c := range n {
		if j+1 == re {
			s[i] = result + string(n[j])
		}
		if isWhiteSpace(n[j]) {
			if i <= x {
				s[i] = result
				i++
				result = ""
			}
		} else {
			result += string(c)
		}
	}
	return s
}

func isWhiteSpace(r byte) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
