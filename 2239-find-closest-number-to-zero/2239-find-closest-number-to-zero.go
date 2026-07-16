func findClosestNumber(nums []int) int {
    num := abs(nums[0])
    ans := nums[0]
    for _, v := range nums {
        if num > abs(v) {
            num = abs(v)
            ans = v
        } else if num == abs(v) && v > 0 {
            ans = v
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}