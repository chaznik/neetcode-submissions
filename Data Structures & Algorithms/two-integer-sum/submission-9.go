func twoSum(nums []int, target int) []int {
    holder := make(map[int]int)
    for i, num := range nums {
        indexAsValue, exists := holder[target - num]
        if exists {
            return []int {indexAsValue, i}
        } else {
            holder[num] = i
        }
    }
    return []int {}
}
