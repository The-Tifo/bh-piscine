package piscine

func Sqrt(num int) int {
	if num < 0 {
		return 0
	}
	root := 0
	for i := 0; i <= num/2+1; i++ {
		if i*i == num {
			root = i
			break
		}
	}
	return root
}
