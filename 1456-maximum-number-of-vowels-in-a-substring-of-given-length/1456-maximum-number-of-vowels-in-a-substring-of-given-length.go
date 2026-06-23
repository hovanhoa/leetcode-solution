func maxVowels(s string, k int) int {
    ans := 0    
    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            ans += 1
        }
    }

    cur := ans
    for i := k; i < len(s); i++ {
        if isVowel(s[i]) {
            cur += 1
        }

        if isVowel(s[i-k]) {
            cur -= 1
        }

        ans = max(ans, cur)
    }

    return ans
}



func isVowel(c byte) bool {
    vowels := []byte{'a', 'e', 'i', 'o', 'u'}
    for _, v := range vowels {
        if v == c {
            return true
        }
    }

    return false
}
