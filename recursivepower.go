package piscine

func RecursivePower(nb int, power int) int {
	if nb != 0 {
		if power > 0 && power < 36 {
			return nb * RecursivePower(nb, power-1)
		} else if power < 0 {
			return 0
		} else if power == 0 {
			return 1
		}
	} else {
		if power != 0 {
			return 0
		} else {
			return 1
		}
	}
	return 1
}
