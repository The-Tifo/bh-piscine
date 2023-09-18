package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Arguments := os.Args
	botato := []rune(Arguments[0])
	for i, t := range botato {
		if i > 1 {
			z01.PrintRune(t)
		}
	}
	z01.PrintRune('\n')
}
