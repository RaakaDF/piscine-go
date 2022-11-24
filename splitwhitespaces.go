package piscine

func SplitWhiteSpaces(n string) []string {
	x := 0
	re := 0
	result := ""
	for index, c := range n {
		if c == ' ' && n[index+1] != ' ' {
			re++
		}
	}
	res := make([]string, re+1)

	for _, r := range n {
		if isSeparator(r) {
			if result != "" {
				res[x] = result
				x++
				result = ""

			}
		} else {
			result += string(r)
		}
	}
	size := 0
	for z := range res {
		size++
		z++
	}
	if result != "" {
		res[size-1] = result
	}
	return res
}

func isSeparator(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
