package piscine

func Rot14(str string) string {
	x := []rune(str)
	var res string

	for i := 0; i < len(x); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			if x[i] >= 'm' {
				x[i] = x[i] - 12
			} else {
				x[i] = x[i] + 14
			}
		} else if x[i] >= 'A' && x[i] <= 'Z' {
			if x[i] >= 'M' {
				x[i] = x[i] - 12
			} else {
				x[i] = x[i] + 14
			}
		}
		res += string(x[i])
	}
	return res
}
