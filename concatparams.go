package piscine

func ConcatParams(args []string) string {
	var t string
	size := 0
	for i := range args {
		i++
		size++
	}
	for i, v := range args {
		t += v
		if i != size-1 {
			t += "\n"
		}
	}
	return t
}
