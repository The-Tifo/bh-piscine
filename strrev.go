package piscine

func StrRev(s string) string {
	a := ""
	arr := []rune(s)

	for i := len(s) - 1; i >= 0; i-- {
		a += string(arr[i])
	}
	return a
}
