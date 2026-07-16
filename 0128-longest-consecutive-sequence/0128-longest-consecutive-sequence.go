func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    m := map[int]bool{}
    for _, v := range nums {
        m[v] = true
    }

    ans := 1
    for v := range m {
        if !m[v-1] {
            cnt := 1
            for m[v+1] {
                cnt += 1
                v += 1
            }

            ans = max(ans, cnt)
        }
    }

    return ans
}