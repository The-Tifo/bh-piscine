package piscine

func IsAlpha(s string) bool {
	for _, t := range s {
		if (t < 'a' || t > 'z') && (t < 'A' || t > 'Z') && (t < '0' || t > '9') {
			return false
		}
	}
	return true
}
