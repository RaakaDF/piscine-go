package piscine

// AdvancedSortWordArr sort a word characters in ascending order
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	lena := 0
	for i := range a {
		lena = i + 1
	}

	for i := 0; i < lena-1; i++ {
		for j := i + 1; j < lena; j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
