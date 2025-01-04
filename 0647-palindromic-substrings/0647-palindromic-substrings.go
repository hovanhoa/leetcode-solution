func countSubstrings(s string) int {
    ans := 0
    for i := range s {
        ans += countPalind(s, i, i) + countPalind(s, i, i+1)
    }

    return ans
}

func countPalind(s string, l, r int) int {
    count := 0

    for l >= 0 && r < len(s) && s[l] == s[r] {
        count += 1
        l -= 1
        r += 1
    }

    return count
}