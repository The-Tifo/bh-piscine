package piscine

func ShoppingListSort(list []string) []string {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if len(list[j]) > len(list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
