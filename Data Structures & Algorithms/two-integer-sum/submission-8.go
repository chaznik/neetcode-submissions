func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        current := nums[i]
        for j := i + 1; j < len(nums); j++ {
            next := nums[j]
            if target - current == next {
                return []int {i, j}
            }
        }
    }
    return []int{}
}
