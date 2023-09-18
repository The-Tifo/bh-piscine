package piscine

func IsLower(s string) bool {
	for _, t := range s {
		if t < 'a' || t > 'z' {
			return false
		}
	}
	return true
}
