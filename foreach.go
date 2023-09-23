package piscine

func ForEach(f func(int), a []int) {
	for _, t := range a {
		f(t)
	}
}
