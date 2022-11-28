package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	length := 0
	for i := range a {
		length = i + 1
	}
	left := false
	right := false
	for i := 0; i < length-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			left = true
		} else if f(a[i], a[i+1]) < 0 {
			right = true
		}
	}
	if left && right {
		return false
	} else {
		return true
	}
}
