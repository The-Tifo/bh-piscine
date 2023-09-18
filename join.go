package piscine

func Join(strs []string, sep string) string {
	z := ""
	for i, t := range strs {
		if i < len(strs)-1 {
			z += t + sep
		} else {
			z += t
		}
	}
	return z
}
