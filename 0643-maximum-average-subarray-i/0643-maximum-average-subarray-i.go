func findMaxAverage(nums []int, k int) float64 {
    s := 0
    for _, v := range nums[:k] {
        s += v
    }

    res := float64(s)/float64(k)

    i := k
    for i < len(nums) {
        s += nums[i] - nums[i-k]
        res = max(res, float64(s)/float64(k))

        i += 1
    }

    return res
}