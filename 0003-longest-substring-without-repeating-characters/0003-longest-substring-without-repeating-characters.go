func lengthOfLongestSubstring(s string) int {
    ans := 0
    set := map[byte]bool{}
    l, r := 0, 0
    for r < len(s) {
        if _, ok := set[s[r]]; !ok {
            set[s[r]] = true
            r++
            ans = max(ans, r - l)
        } else {
            delete(set, s[l])
            l++
        }
    }
    return ans
}
