package piscine

func IsNumeric(s string) bool {
	for _, t := range s {
		if t >= '0' && t <= '9' {
		} else {
			return false
		}
	}
	return true
}
