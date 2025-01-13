func minSwaps(s string) int {
    close, maxClose := 0, 0

    for _, v := range s {
        if v == '[' {
            close -= 1
        } else {
            close += 1
        }

        maxClose = max(maxClose, close)
    }

    return (maxClose + 1) / 2
}