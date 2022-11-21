package piscine

func IterativeFactorial(nb int) int {
	res := nb

	if nb > 25 || nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb >= 1 {
		for i := 1; i < nb; i++ {
			res = res * i
		}
	}
	return res
}
