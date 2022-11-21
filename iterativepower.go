package piscine

func IterativePower(nb int, power int) int {
	res := 1
	if nb != 0 {
		if power > 0 && power < 36 {
			for i := 1; i <= power; i++ {
				res = res * nb
			}
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
	return res
}
