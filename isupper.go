package piscine

func IsUpper(s string) bool {
	for _, t := range s {
		if t < 'A' || t > 'Z' {
			return false
		}
	}
	return true
}
