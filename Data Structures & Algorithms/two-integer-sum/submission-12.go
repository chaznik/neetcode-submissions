func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)

    for i, num := range nums {
        wantedIndex, found := indexMap[target - num]
        if found {
            return []int { wantedIndex, i }
        } else {
            _, isDuplicate := indexMap[num]
            if isDuplicate {
                continue
            }
            indexMap[num] = i
        }
    }
    return []int {}
}