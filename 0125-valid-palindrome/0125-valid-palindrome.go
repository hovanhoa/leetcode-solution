func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r { 
        for l < r && !isAlpha(s[l]) {
            l += 1
        }

        for l < r && !isAlpha(s[r]) {
            r -= 1
        }

        if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
            return false
        }

        l, r = l + 1, r - 1
    }
    
    return true
}

func isAlpha(a byte) bool {
    return unicode.IsLetter(rune(a)) || unicode.IsNumber(rune(a))
}