package piscine

func BasicAtoi(s string) int {
	arr := []rune(s)
	t := 0
	for _, v := range arr {
		t = (t * 10) + int(v-'0')
	}
	return t
}
