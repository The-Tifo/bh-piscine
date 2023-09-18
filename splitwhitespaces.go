package piscine

func SplitWhiteSpaces(s string) []string {
	a := 0
	ln := 0

	ok2 := false
	for c := range s {

		if ok2 && c != 0 && (s[c-1] == '\n' || s[c-1] == '\t' || s[c-1] == ' ') && s[c] != '\n' && s[c] != '\t' && s[c] != ' ' {
			ln++
		}
		if s[c] != '\n' && s[c] != '\t' && s[c] != ' ' {
			ok2 = true
		}
	}
	ln++

	ans := make([]string, ln)
	index := 0
	myS := ""
	for i, w := range s {
		if w == '\n' || w == ' ' || w == '\t' {
			if myS != "" {
				ans[index] = myS
				index++
				myS = ""
				a = i
			}
		} else {
			if w != ' ' {
				myS = myS + string(w)
			}
		}
	}

	if s[a+1:] != "" && s[a+1:] != " " {
		ans[index] = s[a+1:]
	}
	return ans
}
