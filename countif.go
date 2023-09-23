package piscine

func CountIf(f func(string) bool, tab []string) int {
	z := 0
	for _, t := range tab {
		if f(t) {
			z++
		}
	}
	return z
}
