package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Arguments := os.Args
	for t := range Arguments {
		if t > 0 {
			botato := []rune(Arguments[t])
			for z := range botato {
				z01.PrintRune(botato[z])
			}
			z01.PrintRune('\n')
		}
	}
}
