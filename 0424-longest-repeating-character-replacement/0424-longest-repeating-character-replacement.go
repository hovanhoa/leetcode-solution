func characterReplacement(s string, k int) int {
    ans := 0
    m := map[byte]int{}

    l, maxf := 0, 0
    for r := 0; r < len(s); r++ {
        m[s[r]] += 1
        maxf = max(maxf, m[s[r]])
        
        for (r - l + 1) - maxf > k {
            m[s[l]] -= 1
            l += 1
        }

        ans = max(ans, r - l + 1)
    }   

    return ans
}