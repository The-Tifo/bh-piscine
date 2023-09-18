package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	t := 0
	// sign := 1
	for _, v := range arr {
		if v >= '0' && v <= '9' {
			t = (t * 10) + int(v-'0')
		} else {
			return 0
		}
	}
	return t
}
