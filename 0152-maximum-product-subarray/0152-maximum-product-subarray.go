func maxProduct(nums []int) int {
    ans := max(nums...)
    curMin, curMax := 1, 1
    for _, v := range nums {
        if v == 0 {
            curMin, curMax = 1, 1
            continue
        }

        tmp := curMax * v
        curMax = max(curMax * v, curMin * v, v)
        curMin = min(tmp, curMin * v, v)
        ans = max(ans, curMax)
    }

    return ans
}

func max(nums ...int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }

    return max
}

func min(nums ...int) int {
    min := nums[0]
    for _, v := range nums {
        if v < min {
            min = v
        }
    }

    return min
}