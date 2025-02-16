func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    l, r := 0, len(nums) - 1
    for l <= r {
        m := (l + r)/2
        if nums[m] > target {
            r = m - 1
        } else if nums[m] < target {
            l = m + 1
        } else {
            first, last := m, m
            for first > 0 && nums[first-1] == nums[first] {
                first -= 1
            }

            for last < len(nums) - 1 && nums[last+1] == nums[last] {
                last += 1
            }

            res[0], res[1] = first, last
            break
        }
    }

    return res
}