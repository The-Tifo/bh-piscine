package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	upperCase := false
	if args[0] == "--upper" {
		upperCase = true
		args = args[1:]
	}
	for _, arg := range args {
		position := 0
		for _, digit := range arg {
			if digit < '0' || digit > '9' {
				position = -1
				break
			}
			position = position*10 + int(digit-'0')
		}
		if position < 1 || position > 26 {
			z01.PrintRune(' ')
			continue
		}
		letter := 'a' + rune(position) - 1
		if upperCase {
			letter -= 32
		}
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
