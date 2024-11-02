func lengthOfLastWord(s string) int {
    i, ans := len(s) - 1, 0
    for s[i] == ' ' {
        i -= 1
    }

    for i >= 0 && s[i] != ' ' {
        ans += 1
        i -= 1
    }

    return ans
}