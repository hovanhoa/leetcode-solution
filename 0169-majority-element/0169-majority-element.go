func majorityElement(nums []int) int {
    ans, max, m := 0, 0, map[int]int{}
    for _, v := range nums {
        m[v] += 1
        if m[v] > max {
            max = m[v]
            ans = v
        }
    }

    return ans
}