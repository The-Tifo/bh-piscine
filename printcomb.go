package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for l := i + 1; l <= '8'; l++ {
			for f := l + 1; f <= '9'; f++ {
				z01.PrintRune(i)
				z01.PrintRune(l)
				z01.PrintRune(f)
				if i == '7' && l == '8' && f == '9' {
					break
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
