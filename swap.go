package piscine

func Swap(a *int, b *int) {
	c := *a
	y := *b
	*a = y
	*b = c
}
