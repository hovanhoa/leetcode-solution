func maxSubArray(nums []int) int {
    ans := nums[0]
    curSum := 0
    for _, v := range nums {
        if curSum < 0 {
            curSum = 0
        }

        curSum += v
        if ans < curSum {
            ans = curSum
        }
    }

    return ans
}