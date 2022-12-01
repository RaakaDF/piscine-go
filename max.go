package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	myMax := a[0]
	for _, v := range a {
		if v > myMax {
			myMax = v
		}
	}
	return myMax
}
