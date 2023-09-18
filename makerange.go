package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		arr := make([]int, max-min)
		for i := 0; i < (max - min); i++ {
			arr[i] = min + i
		}
		return arr
	}
}
