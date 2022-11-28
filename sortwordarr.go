package piscine

// SortWordArr function sorts a string array by ascii  in asc order.
func SortWordArr(a []string) {
	lena := 0
	for i := range a {
		lena = i + 1
	}
	for i := 0; i < lena-1; i++ {
		for j := i + 1; j < lena; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
