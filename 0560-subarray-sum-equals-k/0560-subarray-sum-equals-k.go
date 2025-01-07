func subarraySum(nums []int, k int) int {
    ans := 0
    curSum := 0
    m := map[int]int{0:1}

    for _, v := range nums {
        curSum += v
        diff := curSum - k

        ans += m[diff]
        m[curSum]+= 1
    }

    return ans
}