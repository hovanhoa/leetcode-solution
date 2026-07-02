func singleNumber(nums []int) int {
    n := nums[0]
    for i := 1; i < len(nums); i++ {
        n ^= nums[i]
    }

    return n
}