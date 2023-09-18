package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for l := '0'; l <= '9'; l++ {
			for f := '0'; f <= '9'; f++ {
				for t := '0'; t <= '9'; t++ {
					if i < f || (i == f && l < t) {
						z01.PrintRune(i)
						z01.PrintRune(l)
						z01.PrintRune(' ')
						z01.PrintRune(f)
						z01.PrintRune(t)
						if i == '9' && l == '8' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
