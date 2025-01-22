func lengthOfLongestSubstring(s string) int {
    ans := 0
    m := map[byte]bool{}
    l := 0
    for r := 0; r < len(s); r++ {
        for m[s[r]] {
            m[s[l]] = false
            l += 1
        }

        m[s[r]] = true
        ans = max(ans, r - l + 1)
    }

    return ans
}