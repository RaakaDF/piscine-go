package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb <= 1 {
			return 2
		}
		for n := 2; n*n <= nb; n++ {
			if nb%n == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb

	} else {
		return 2
	}
}
