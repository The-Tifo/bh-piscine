package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	result := 1
	for i := power; i > 0; i-- {
		result *= nb
	}
	return result
}
