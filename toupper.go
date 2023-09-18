package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for i, z := range arr {
		if arr[i] >= 97 && arr[i] <= 122 {
			arr[i] = z - 32
		}
	}
	return string(arr)
}
