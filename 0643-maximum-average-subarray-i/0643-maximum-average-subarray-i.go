func findMaxAverage(nums []int, k int) float64 {
    ans := 0
    for i := 0; i < k; i++ {
        ans += nums[i]
    }

    cur := ans
    for i := k; i < len(nums); i++ {
        cur += nums[i] - nums[i-k]
        ans = max(ans, cur)
    }

    return float64(ans)/float64(k)
}
