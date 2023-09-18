package piscine

func Atoi(s string) int {
	arr := []rune(s)
	t := 0
	sign := 1
	for i, v := range arr {
		if v >= '0' && v <= '9' {
			t = (t * 10) + int(v-'0')
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v == '-' && i == 0 {
			sign = -1
		} else {
			return 0
		}
	}
	return t * sign
}
