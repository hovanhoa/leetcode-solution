func characterReplacement(s string, k int) int {
    count := map[byte]int{}
    ans := 0

    l, maxf := 0, 0
    for r := 0; r < len(s); r++ {
        count[s[r]] += 1
        maxf = max(maxf, count[s[r]])

        for (r - l + 1) - maxf > k {
            count[s[l]] -= 1
            l += 1
        }

        ans = max(ans, r - l + 1)
    }

    return ans
}