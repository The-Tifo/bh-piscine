package piscine

func BasicJoin(elems []string) string {
	z := ""
	for _, elem := range elems {
		z += elem
	}
	return z
}
