func findMaxAverage(nums []int, k int) float64 {
    result := 0
    for i := 0; i < k; i ++ {
        result += nums[i]
    }

    maxVal := result
    for i := k; i < len(nums); i ++ {
        maxVal = maxVal + nums[i] - nums[i-k]
        if maxVal > result {
            result = maxVal
        }
    }

    return float64(result)/float64(k)
}