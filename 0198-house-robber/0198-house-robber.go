func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    for i := range nums {
        if i < 2 {
            continue
        } else if i == 2 {
            nums[i] = nums[i] + nums[i-2]
            continue
        }

        nums[i] += max(nums[i-2], nums[i-3])
    }

    return max(nums[len(nums)-1], nums[len(nums)-2])
}