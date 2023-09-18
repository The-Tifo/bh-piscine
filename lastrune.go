package piscine

func LastRune(s string) rune {
	arr := []rune(s)
	ben := len(arr)
	return arr[ben-1]
}
