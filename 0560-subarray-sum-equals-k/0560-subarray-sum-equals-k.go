func subarraySum(nums []int, k int) int {
    // map diff to count
    m := map[int]int{0: 1}
    curSum := 0
    ans := 0

    for _, v := range nums {
        curSum += v
        diff := curSum - k

        if cnt, ok := m[diff]; ok {
            ans += cnt
        }

        m[curSum] += 1
    }

    return ans
}
