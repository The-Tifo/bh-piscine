package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	ben := len(arr)
	if n > ben || n < 0 || n-1 < 0 {
		return 0
	}
	return arr[n-1]
}
