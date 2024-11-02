func lengthOfLastWord(s string) int {
    isStart := false
    ans := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            ans += 1
            isStart = true
        } else if isStart {
            return ans
        }
    }

    return ans
}