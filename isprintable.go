package piscine

func IsPrintable(s string) bool {
	for _, t := range s {
		if t >= 32 {
		} else {
			return false
		}
	}
	return true
}
