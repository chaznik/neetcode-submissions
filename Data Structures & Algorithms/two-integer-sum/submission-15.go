func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
	var keyFound int

	for i := 0; i < len(nums); i++ {
		keyFound = target - nums[i]
		index, found := numMap[keyFound]
		if found {
			return []int{index, i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}
