func lengthOfLongestSubstring(s string) int {
    ans := 0
    count := make([]int, 256)

    for l, r := 0, 0; r < len(s); r++ {
        count[s[r]] += 1
        for count[s[r]] > 1 {
            count[s[l]] -= 1
            l += 1
        }

        if r - l + 1 > ans {
            ans = r - l + 1
        }
    }

    return ans
}