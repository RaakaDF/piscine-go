package piscine

func BasicAtoi(s string) int {
	byteRevStr := []byte(s)
	lengthStr := -1

	for range s {
		lengthStr++
	}

	for index := range s {
		byteRevStr[lengthStr-index] = s[index]
		byteRevStr[lengthStr-index] -= 48
	}

	num := 0
	tenth := 1
	for index := range s {
		num = num + (int(byteRevStr[index]) * tenth)
		tenth *= 10
	}
	return num
}
