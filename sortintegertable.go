package piscine

func SortIntegerTable(table []int) {
	p := 0
	for t := 1; t < len(table); t++ {
		for b := 1; b < len(table); b++ {
			if table[b-1] > table[b] {
				p = table[b]
				table[b] = table[b-1]
				table[b-1] = p
			}
		}
	}
}
