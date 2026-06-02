func findMaxAverage(nums []int, k int) float64 {
    s := 0
    for i := 0; i < k; i ++ {
        s += nums[i]
    }

    maxSum := s
    for i := k; i < len(nums); i++ {
        s += nums[i] - nums[i-k]
        maxSum = max(maxSum, s)
    }

    return float64(maxSum)/float64(k)
}