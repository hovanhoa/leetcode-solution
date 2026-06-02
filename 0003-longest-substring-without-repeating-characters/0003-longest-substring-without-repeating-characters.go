func lengthOfLongestSubstring(s string) int {
    ans := 0
    m := map[byte]bool{}
    l, r := 0, 0
    for r < len(s) {
        if _, ok := m[s[r]]; ok {
            delete(m, s[l])
            l += 1
        } else {
            m[s[r]] = true
            r += 1
            ans = max(ans, r - l)
        }
    }

    return ans
}