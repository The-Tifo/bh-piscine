package piscine

func FindNextPrime(num int) int {
	if num < 2 {
		return 2
	}
	if num == 2 {
		return num
	}
	for i := num; ; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return i
		}
	}
}
