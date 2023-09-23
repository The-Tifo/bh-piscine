package main

import (
	"fmt"
	"os"
	"strconv"
)

func convertIntToRune(nbr int) string {
	return strconv.Itoa(nbr)
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		return
	}

	sign := 0
	switch arguments[1] {
	case "+":
		sign = 0
	case "-":
		sign = 1
	case "*":
		sign = 2
	case "/":
		sign = 3
	case "%":
		sign = 4
	default:
		return
	}

	firstNbr, err := strconv.Atoi(arguments[0])
	if err != nil {
		return
	}

	secondNbr, err := strconv.Atoi(arguments[2])
	if err != nil {
		return
	}

	if secondNbr == 0 && (arguments[1] == "/" || arguments[1] == "%") {
		errorMsg := "No division by 0"
		if arguments[1] == "%" {
			errorMsg = "No modulo by 0"
		}
		fmt.Fprintln(os.Stderr, errorMsg)
		return
	}

	arrayOfFunctions := []func(int, int) int{plus, minus, times, div, mod}
	result := apply(arrayOfFunctions[sign], firstNbr, secondNbr)

	if result >= 2147483648 || result <= -2147483648 {
		return
	}

	fmt.Fprintln(os.Stderr, convertIntToRune(result))
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func apply(f func(int, int) int, a int, b int) int {
	return f(a, b)
}
