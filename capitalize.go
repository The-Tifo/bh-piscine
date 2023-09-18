package piscine

func Capitalize(s string) string {
	count := 1
	after := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			if count == 1 {
				after += string(s[i])
			} else {
				after += string(s[i] + 32)
			}
			count++
		} else if s[i] >= 97 && s[i] <= 122 {
			if count == 1 {
				after += string(s[i] - 32)
			} else {
				after += string(s[i])
			}
			count++
		} else if s[i] >= 48 && s[i] <= 57 {
			after += string(s[i])
			count++
		} else {
			count = 1
			after += string(s[i])
		}
	}
	return after
}
