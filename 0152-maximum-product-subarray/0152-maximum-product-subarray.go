func maxProduct(nums []int) int {
    ans := slices.Max(nums)
    curMin, curMax := 1, 1
    for _, v := range nums {
        if v == 0 {
            curMin, curMax = 1, 1
            continue
        }

        tmp := curMin
        curMin = min(curMin * v, curMax * v, v)
        curMax = max(tmp * v, curMax * v, v)
        ans = max(curMax, ans)        
    }

    return ans
}