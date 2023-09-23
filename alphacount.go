package piscine

func AlphaCount(s string) int {
	count := 0
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 65 && arr[i] <= 90 {
			count++
		}
		if arr[i] >= 97 && arr[i] <= 122 {
			count++
		}
	}
	return count
}
