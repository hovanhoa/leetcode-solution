func subarraySum(nums []int, k int) int {
    ans := 0
    m := map[int]int{0: 1}
    s := 0
    for _, v := range nums {
        s += v
        diff := s - k

        if _, ok := m[diff]; ok {
            ans += m[diff]
        }

        m[s] += 1
    }

    return ans
}