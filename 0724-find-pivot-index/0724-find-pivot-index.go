func pivotIndex(nums []int) int {
    rSum := 0
    for _, v := range nums {
        rSum += v
    }

    lSum := 0
    for i, v := range nums {
        rSum -= v
        if i > 0 {
            lSum += nums[i-1]
        }

        if rSum == lSum {
            return i
        }
    }

    return -1
}