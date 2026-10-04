func hasDuplicate(nums []int) bool {
    numsMap := make(map[int] int)

	for i := 0; i < len(nums); i++ {
		_, exists := numsMap[nums[i]]
		if exists {
			return true
		} else {
			numsMap[nums[i]] = i
		}
	}
	return false
}
