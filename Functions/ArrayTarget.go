package reverse

func INTinArray(ArrayInt []int, TargetINt int) bool {
	for _, num := range ArrayInt {
		if num == TargetINt {
			return true
		}
	}
	return false
}
