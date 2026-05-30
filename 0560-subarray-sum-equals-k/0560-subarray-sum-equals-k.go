func subarraySum(nums []int, k int) int {
    m := map[int]int{0: 1}
    ans := 0
    s := 0
    for _, v := range nums {
        s += v
        diff := s - k
        if cnt, ok := m[diff]; ok {
            ans += cnt
        }

        m[s] += 1
    }

    return ans
}