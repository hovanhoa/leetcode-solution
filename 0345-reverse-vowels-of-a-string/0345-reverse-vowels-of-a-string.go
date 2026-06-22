func reverseVowels(s string) string {
    l, r := 0, len(s) - 1 
    ans := []byte(s)
    for l < r {
        for l < r && !isVowel(ans[l]) {
            l += 1
        }

        for l < r && !isVowel(ans[r]) {
            r -= 1
        }

        ans[l], ans[r] = ans[r], ans[l]
        l, r = l + 1, r - 1
    }

    return string(ans)
}

func isVowel(s byte) bool {
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for _, v := range vowels {
        if s == v {
            return true
        }
    }

    return false
}