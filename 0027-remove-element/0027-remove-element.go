func removeElement(nums []int, val int) int {
    res := 0
    l, r := 0, 0
    for r < len(nums) {
        if nums[r] != val {
            nums[l] = nums[r]
            l += 1
            res += 1
        }

        r += 1
    }

    return res
}