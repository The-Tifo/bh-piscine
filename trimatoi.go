package piscine

func TrimAtoi(s string) int {
	var t int = 1
	var a int = 0
	var z bool = false
	for _, q := range s {
		if q == '-' && !z {
			t = -1
		} else if q >= '0' && q <= '9' {
			a = a*10 + int(q-'0')
			z = true
		}
	}
	return a * t
}
