func hasDuplicate(nums []int) bool {
    numDict := make(map[int]int)

    for _, num := range nums {
        _, exists := numDict[num]
        if exists {
            return true
        } else {
            numDict[num] = 1
        }
    }
    return false;
}
