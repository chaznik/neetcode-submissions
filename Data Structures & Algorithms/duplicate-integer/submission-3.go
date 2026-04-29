func hasDuplicate(nums []int) bool {
    holder := make(map[int]int)

	for i, num := range nums {
		_, exists := holder[num]
		if exists {
			return true
		} else {
			holder[num] = i
		}
	}
	return false
}
