func isPalindrome(s string) bool {
    res := strings.ToLower(s)
    str := []byte{}
    for i := range res {
        if isAlphaByte(res[i]) {
            str = append(str, res[i])
        }
    }

    l, r := 0, len(str) - 1
    for l < r {
        if str[l] == str[r] {
            l += 1
            r -= 1
        } else {
            return false
        }
    }

    return true
}

func isAlphaByte(b byte) bool {
    return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
