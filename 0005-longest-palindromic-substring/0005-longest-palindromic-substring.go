func longestPalindrome(s string) string {
    ans := ""
    for i := 0; i < len(s); i++ {
        l, r := i, i
        for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
            if len(ans) < r - l + 1 {
                ans = s[l:r+1]
            }

            l, r = l - 1, r + 1
        }

        l, r = i, i + 1
        for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
            if len(ans) < r - l + 1 {
                ans = s[l:r+1]
            }

            l, r = l - 1, r + 1
        }
    }

    return ans
}