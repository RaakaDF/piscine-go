package piscine

func Compact(ptr *[]string) int {
	len := 0
	for _, a := range *ptr {
		if a != "" {
			len++
		}
	}
	count12 := 0
	arr_new := make([]string, len)
	for _, a := range *ptr {
		if a != "" {
			arr_new[count12] = a
			count12++
		}
	}
	*ptr = arr_new
	return len
}
