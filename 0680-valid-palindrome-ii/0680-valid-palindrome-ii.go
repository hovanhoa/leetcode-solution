func validPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return isRevert(s, l + 1, r) || isRevert(s, l, r - 1)
        }

        l, r = l + 1, r - 1
    }

    return true
}

func isRevert(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }

        l, r = l + 1, r - 1
    }

    return true
}