package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for i, z := range arr {
		if arr[i] >= 65 && arr[i] <= 90 {
			arr[i] = z + 32
		}
	}
	return string(arr)
}
