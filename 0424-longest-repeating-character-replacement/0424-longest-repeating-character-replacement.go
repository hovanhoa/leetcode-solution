func characterReplacement(s string, k int) int {
    ans := 0
    m := map[byte]int{}

    l, maxM := 0, 0
    for r := 0; r < len(s); r++ {
        m[s[r]] += 1
        maxM = max(maxM, m[s[r]])

        for (r - l + 1) - maxM > k {
            m[s[l]] -= 1
            l += 1
        }

        ans = max(ans, r - l + 1)
    }

    return ans
}