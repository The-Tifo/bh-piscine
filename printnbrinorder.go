package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var d [10]int
	for n > 0 {
		d[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < d[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
