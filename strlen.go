package piscine

func StrLen(s string) int {
	arr := []rune(s)
	a := 0
	for t := range arr {
		a = t + 1
	}
	return a
}
