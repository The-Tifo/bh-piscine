package piscine

func Index(s string, toFind string) int {
	slen := len(s)
	tlen := len(toFind)
	if tlen == 0 {
		return 0
	}
	for i := 0; i <= slen-tlen; i++ {
		if s[i:i+tlen] == toFind {
			return i
		}
	}
	return -1
}
