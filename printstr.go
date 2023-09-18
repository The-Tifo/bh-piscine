package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	arr := []rune(s)
	for _, t := range arr {
		z01.PrintRune(t)
	}
}
