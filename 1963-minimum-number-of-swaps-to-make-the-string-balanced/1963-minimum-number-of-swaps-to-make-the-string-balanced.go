func minSwaps(s string) int {
    ans, cur := 0, 0
    for _, v := range s {
        if v == ']' {
            cur += 1
        } else {
            cur -= 1
        }

        ans = max(ans, cur)
    }

    return (ans + 1)/2
}