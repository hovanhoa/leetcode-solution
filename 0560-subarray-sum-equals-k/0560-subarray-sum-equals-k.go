func subarraySum(nums []int, k int) int {
    m := map[int]int{0: 1}
    s := 0
    res := 0
    for _, v := range nums {
        s += v
        if _, ok := m[s-k]; ok {
            res += m[s-k]
        }

        m[s] += 1
    }

    return res
}