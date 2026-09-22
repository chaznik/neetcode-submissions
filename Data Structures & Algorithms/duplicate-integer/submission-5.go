func hasDuplicate(nums []int) bool {
    numDict := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        valueAtKey := numDict[nums[i]]
        if valueAtKey != 0 {
            return true;
        }
        numDict[nums[i]] = 1
    }
    return false
}
