package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Arguments := os.Args[1:]
	for i := range Arguments {
		for j := i + 1; j < len(Arguments); j++ {
			if Arguments[i] > Arguments[j] {
				Arguments[i], Arguments[j] = Arguments[j], Arguments[i]
			}
		}
	}
	for t := range Arguments {
		botato := []rune(Arguments[t])
		for z := range botato {
			z01.PrintRune(botato[z])
		}
		z01.PrintRune('\n')
	}
}
